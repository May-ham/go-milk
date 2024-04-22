package controllers

import (
	"go-tavern/tokenization"
	"net/http"

	"github.com/gorilla/mux"
)

func CountTokens(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	modelName := vars["name"]
	switch modelName {
	case "ai21":
		tokenization.TokenizeAI21(w, r)
	case "palm":
		tokenization.TokenizePalm(w, r)
	default:

	}
}
