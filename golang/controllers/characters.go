package controllers

import (
	"fmt"
	"go-tavern/imageutils"
	"go-tavern/paths"
	_ "image/jpeg"
	_ "image/png"
	"io"
	"log"
	"net/http"
	"os"
	"path"

	"github.com/gorilla/mux"
)

func GetAllCharacters(w http.ResponseWriter, r *http.Request) {
	files, err := os.ReadDir(paths.CharactersPath)
	if err != nil {
		http.Error(w, "Unable to read directory", http.StatusInternalServerError)
		return
	}

	for _, file := range files {
		fmt.Fprintln(w, file.Name())
	}
}

func CreateCharacter(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(10 << 20) // Limit your maxMultipartMemory
	file, handler, err := r.FormFile("myFile")
	if err != nil {
		fmt.Println("Error Retrieving the File")
		fmt.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer file.Close()

	err = imageutils.ValidatePNG(file)
	if err != nil {
		fmt.Println("File is not a character card")
		fmt.Println(err)
		http.Error(w, "File is not a character card", http.StatusInternalServerError)
		return
	}

	filename := handler.Filename

	err = checkFileExists(paths.CharactersPath + filename)

	log.Println(filename)
	log.Println(err)

	if err != nil {
		filename = generateNewFilename(paths.CharactersPath + filename)
	}

	log.Println(filename)

	dst, err := os.Create(paths.CharactersPath + filename)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Println(1, err.Error())
		return
	}
	defer dst.Close()

	if _, err := io.Copy(dst, file); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Println(2, err.Error())
		return
	}

	fmt.Fprintf(w, "Successfully Uploaded File\n")
}

func checkFileExists(filePath string) error {
	_, err := os.Stat(filePath)
	if err == nil {
		return fmt.Errorf("file %q already exists", filePath)
	} else if os.IsNotExist(err) {
		return nil
	} else {
		// File may or may not exist. See the underlying error for more information.
		return err
	}
}

func generateNewFilename(filePath string) string {
	dir := path.Dir(filePath) // Extract the directory from the provided path
	ext := path.Ext(filePath)
	name := path.Base(filePath)
	base := name[:len(name)-len(ext)]

	for i := 1; ; i++ {
		newFilename := fmt.Sprintf("%s (%d)%s", base, i, ext)
		newPath := path.Join(dir, newFilename)
		if checkFileExists(newPath) != nil {
			continue
		}
		return newFilename
	}
}

func DownloadCharacter(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fileName := vars["name"]
	file, err := os.Open(paths.CharactersPath + fileName)
	if err != nil {
		http.Error(w, "File not found", 404)
		return
	}
	defer file.Close()
	io.Copy(w, file)
}
