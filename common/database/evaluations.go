package database

import (
	"log"

	"github.com/Mardiniii/serapis/common/models"
	"github.com/lib/pq"
	"github.com/pkg/errors"
)

func (conn *Postgres) createEvaluationsTable() (err error) {
	if _, err = conn.Db.Exec(evaluationsTable); err != nil {
		err = errors.Wrapf(err, "Can not create evaluations table (%s)", evaluationsTable)
		return
	}

	return
}

// CreateEvaluation adds a new evaluation record to the database
func (conn *Postgres) CreateEvaluation(eval *models.Evaluation) (err error) {
	err = conn.Db.QueryRow(createEvaluation,
		eval.UserID,
		eval.Status,
		eval.Language,
		eval.Code,
		pq.Array(eval.Stdin),
		eval.Dependencies,
		eval.Git,
	).Scan(&eval.ID, &eval.CreatedAt)
	if err != nil {
		log.Println(err)
	}
	return
}
