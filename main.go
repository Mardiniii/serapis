package main

import (
	"github.com/Mardiniii/serapis/api"
	"github.com/Mardiniii/serapis/evaluator"
)

func main() {
	// Ruby Evaluation
	evaluator.Evaluate("ruby", "puts \"Hola Mundo vamos con toda\"")

	// Node Evaluation
	evaluator.Evaluate("node", "console.log('hello world');")

	api.Init()
}
