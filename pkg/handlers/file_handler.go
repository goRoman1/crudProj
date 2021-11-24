package handlers

import (
	"crudProj/internal/file_uploader"
	"fmt"
	"io/ioutil"
	"net/http"
)

func UploadFile(w http.ResponseWriter, r *http.Request){
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

	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println(err)
	}
	tempFile, err := ioutil.TempFile("./../internal/file_uploader/temp-files", "upload-*.сsv")
	if err != nil {
		fmt.Println(err)
	}

	defer tempFile.Close()
//	defer os.Remove(tempFile.Name())
	tempFile.Write(fileBytes)

	file_uploader.Pars(tempFile.Name())


	fmt.Fprintf(w, "Successfully Uploaded File\n")
}

