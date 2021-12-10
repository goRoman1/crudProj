package handlers

import (
	"crudProj/entities"
	"crudProj/pkg/services"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type FileHandler struct {
	fileService services.FileServiceI
}

func NewFileHandler(fileService services.FileServiceI) *FileHandler {
	return &FileHandler{
		fileService: fileService,
	}
}

type FileHandlerI interface {
	UploadFil(w http.ResponseWriter, r *http.Request)
	Test(w http.ResponseWriter, r *http.Request)
}

func (f FileHandler) Test(w http.ResponseWriter, r *http.Request) {
	var scooter entities.Test
	err := json.NewDecoder(r.Body).Decode(&scooter)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotAcceptable)
		return
	}
	err = f.fileService.TestService(&scooter)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotAcceptable)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (f FileHandler)UploadFile(w http.ResponseWriter, r *http.Request){
	fmt.Println("File Upload Endpoint Hit")

	r.ParseMultipartForm(10 << 20)
	file, handler, err := r.FormFile("myFile")
	if err != nil {
		fmt.Println("Error Retrieving the File")
		fmt.Println(err)
		panic(err)
	}
	defer file.Close()
	fmt.Printf("Uploaded File: %+v\n", handler.Filename)
	fmt.Printf("File Size: %+v\n", handler.Size)
	fmt.Printf("MIME Header: %+v\n", handler.Header)

	filePath := f.fileService.InsertScootersToDb(file)
	defer os.Remove(filePath)

	fmt.Fprintf(w, "Successfully Uploaded File\n")
}

