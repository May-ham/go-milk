package controllers

import (
	"encoding/json"
	"go-tavern/completion"
	"net/http"

	"github.com/gorilla/mux"
)

func Completion(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	modelName := vars["name"]
	switch modelName {
	case "ai21":
		var p completion.ParamsAI21

		// Decode the request body into the struct
		err := json.NewDecoder(r.Body).Decode(&p)
		if err != nil {
			// handle error
			return
		}

		completion.Complete(p)
	case "openai":
		var p completion.ParamsOpenAI

		// Decode the request body into the struct
		err := json.NewDecoder(r.Body).Decode(&p)
		if err != nil {
			// handle error
			return
		}

		completion.Complete(p)

	case "palm":
		var p completion.ParamsPalm

		// Decode the request body into the struct
		err := json.NewDecoder(r.Body).Decode(&p)
		if err != nil {
			// handle error
			return
		}

		completion.Complete(p)
	// case "palm":
	// 	tokenizePalm(w, r)
	default:

	}
}
