package controllers

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/Mardiniii/serapis/api/models"
	"github.com/Mardiniii/serapis/evaluator"
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

	// Parse JSON data with User struct
	err = json.Unmarshal(body, &eval)
	if err != nil {
		RespondWithError(w, http.StatusUnprocessableEntity, "The evaluation was not created")
		return
	}
	eval.Language = lang
	exitCode := evaluator.Evaluate(lang, eval.Code)
	eval.ExitCode = exitCode

	RespondWithJSON(w, http.StatusOK, eval)
}
