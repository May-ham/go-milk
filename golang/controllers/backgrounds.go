package controllers

import (
	"fmt"
	"go-tavern/paths"
	"io"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func CreateBackground(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(10 << 20) // Limit your maxMultipartMemory
	file, handler, err := r.FormFile("myFile")
	if err != nil {
		fmt.Println("Error Retrieving the File")
		fmt.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer file.Close()

	dst, err := os.Create(paths.BackgroundsPath + handler.Filename)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer dst.Close()

	if _, err := io.Copy(dst, file); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "Successfully Uploaded File\n")
}

func GetAllBackgrounds(w http.ResponseWriter, r *http.Request) {
	files, err := os.ReadDir(paths.BackgroundsPath)
	if err != nil {
		http.Error(w, "Unable to read directory", http.StatusInternalServerError)
		return
	}

	for _, file := range files {
		fmt.Fprintln(w, file.Name())
	}
}

func DownloadBackground(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fileName := vars["name"]
	file, err := os.Open(paths.BackgroundsPath + fileName)
	if err != nil {
		http.Error(w, "File not found", 404)
		return
	}
	defer file.Close()
	io.Copy(w, file)
}
