package database

import (
	"database/sql"
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
	dep, _ := eval.Dependencies.MarshalJSON()
	git, _ := eval.Git.MarshalJSON()

	err = conn.Db.QueryRow(createEvaluation,
		eval.UserID,
		eval.Status,
		eval.Language,
		eval.Code,
		pq.Array(eval.Stdin),
		dep,
		git,
	).Scan(&eval.ID, &eval.CreatedAt)
	if err != nil {
		log.Println(err)
	}
	return
}

// FindEvaluationByID returns the evaluation with the given id
func (conn *Postgres) FindEvaluationByID(id int) (eval models.Evaluation, err error) {
	var output sql.NullString
	var exitCode sql.NullInt64
	row := conn.Db.QueryRow(evaluationByID, id)

	err = row.Scan(
		&eval.ID,
		&eval.UserID,
		&eval.Status,
		&eval.Language,
		&eval.Code,
		pq.Array(&eval.Stdin),
		&eval.Dependencies,
		&eval.Git,
		&output,
		&exitCode,
		&eval.CreatedAt,
	)
	if output.Valid {
		eval.Output = output.String
	} else {
		eval.Output = ""
	}

	if exitCode.Valid {
		eval.ExitCode = int(exitCode.Int64)
	} else {
		eval.ExitCode = -1
	}
	return
}
