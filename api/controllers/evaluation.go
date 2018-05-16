package controllers

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/Mardiniii/serapis/api/broker"
	db "github.com/Mardiniii/serapis/common/database"
	"github.com/Mardiniii/serapis/common/models"
	"github.com/gorilla/mux"
)

var langs = []string{"ruby", "node"}

func supportedLanguage(lang string) bool {
	for _, e := range langs {
		if e == lang {
			return true
		}
	}
	return false
}

// CreateEvaluation handler to process a new evaluation request
func CreateEvaluation(w http.ResponseWriter, r *http.Request) {
	var eval models.Evaluation
	var statusCode = http.StatusOK

	lang := mux.Vars(r)["language"]
	if !supportedLanguage(lang) {
		RespondWithError(w, http.StatusBadRequest, "Given language is not supported")
		return
	}

	// Extract JSON payload
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		RespondWithError(w, http.StatusBadRequest, "Invalid payload")
		return
	}
	defer r.Body.Close()

	err = json.Unmarshal(body, &eval)
	if err != nil {
		RespondWithError(w, http.StatusUnprocessableEntity, "The evaluation was not created"+err.Error())
		return
	}

	// Parse JSON data with Evaluation struct
	ctx := r.Context()
	user := ctx.Value("user").(models.User)

	eval.UserID = user.ID
	eval.Status = "created"
	eval.Language = lang
	// Create the new evaluation
	err = db.RepoCreateEvaluation(&eval)
	if err != nil {
		RespondWithError(w, http.StatusConflict, "The evaluation was not created")
		return
	}

	// Publish evaluation to the message queue
	resp, err := broker.EvaluationRPC(eval.ID)
	if err != nil {
		RespondWithError(w, http.StatusBadRequest, "Invalid payload")
		return
	}

	id, _ := strconv.Atoi(string(resp))
	eval, err = db.RepoFindEvaluationByID(id)
	if err != nil {
		RespondWithError(w, http.StatusBadRequest, "Can't find processed evaluation")
		return
	}

	if eval.ExitCode != 0 {
		statusCode = http.StatusUnprocessableEntity
	}
	RespondWithJSON(w, statusCode, eval)
}
