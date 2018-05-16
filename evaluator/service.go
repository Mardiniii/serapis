package evaluator

import (
	"fmt"
	"log"
	"strconv"

	db "github.com/Mardiniii/serapis/common/database"
	"github.com/streadway/amqp"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
		panic(fmt.Sprintf("%s: %s", msg, err))
	}
}

// Init starts Evaluator service
func Init() {
	// Create connection
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	// Create channel
	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	// Create queue
	q, err := ch.QueueDeclare(
		"rpc_queue", // name
		false,       // durable
		false,       // delete when usused
		false,       // exclusive
		false,       // no-wait
		nil,         // arguments
	)
	failOnError(err, "Failed to declare a queue")

	// Fair dispatch based on acknowledgment
	err = ch.Qos(
		1,     // prefetch count
		0,     // prefetch size
		false, // global
	)
	failOnError(err, "Failed to set QoS")

	// Queue consumer
	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		false,  // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	failOnError(err, "Failed to register a consumer")

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			// Update eval status
			id, _ := strconv.Atoi(string(d.Body))
			fmt.Println("***Processing evaluation with id:", id)

			err = db.RepoUpdateEvalStatus(id, "evaluating")
			failOnError(err, "Failed to update evaluation status")
			eval, err := db.RepoFindEvaluationByID(id)
			failOnError(err, "Failed to find the evalution")

			// Run evaluation
			exec := executor{&eval}
			exitCode, output, err := exec.Start()
			failOnError(err, "Failed to process evaluation")

			// Update evaluation
			err = db.RepoUpdateEvalResult(id, exitCode, output, "finished")
			failOnError(err, "Failed to update evaluation result")

			err = ch.Publish(
				"",        // exchange
				d.ReplyTo, // routing key
				false,     // mandatory
				false,     // immediate
				amqp.Publishing{
					ContentType:   "application/json",
					CorrelationId: d.CorrelationId,
					Body:          d.Body,
				})
			failOnError(err, "Failed to publish a message")

			// Send acknowledgment to the queue
			d.Ack(false)
		}
	}()

	log.Printf(" [*] Awaiting RPC requests")
	<-forever
}
