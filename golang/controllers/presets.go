package controllers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

type Preset struct {
	Name   string      `json:"presetName"`
	Family string      `json:"presetFamily"`
	Model  string      `json:"presetModel"`
	Params interface{} `json:"presetParams"`
}

func CreatePreset(w http.ResponseWriter, r *http.Request) {
	// Read the request body
	body, err := io.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusInternalServerError)
		return
	}

	// Unmarshal the JSON body into the struct
	var preset Preset
	err = json.Unmarshal(body, &preset)
	if err != nil {
		http.Error(w, "Error unmarshalling request body", http.StatusBadRequest)
		return
	}

	switch preset.Name {
	case "openai":

	}
}

func GetAllPresets(w http.ResponseWriter, r *http.Request) {
	files, err := os.ReadDir("./presets")
	if err != nil {
		http.Error(w, "Unable to read directory", http.StatusInternalServerError)
		return
	}

	for _, file := range files {
		fmt.Fprintln(w, file.Name())
	}
}

func DownloadPreset(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fileName := vars["name"]
	file, err := os.Open("./presets/" + fileName)
	if err != nil {
		http.Error(w, "File not found", 404)
		return
	}
	defer file.Close()
	io.Copy(w, file)
}
