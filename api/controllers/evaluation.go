package controllers

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/Mardiniii/serapis/api/broker"
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

	// Parse JSON data with Evaluation struct
	eval.Language = lang
	err = json.Unmarshal(body, &eval)
	if err != nil {
		RespondWithError(w, http.StatusUnprocessableEntity, "The evaluation was not created"+err.Error())
		return
	}

	// Evaluate request
	// json, err := json.Marshal(eval)
	resp, err := broker.EvaluationRPC("Este es el texto")
	if err != nil {
		RespondWithError(w, http.StatusBadRequest, "Invalid payload")
		return
	}
	fmt.Println(resp)

	if eval.ExitCode != 0 {
		statusCode = http.StatusUnprocessableEntity
	}

	RespondWithJSON(w, statusCode, eval)
}
