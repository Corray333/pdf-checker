package server

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func LoadPDF(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(10 << 20)

	// Get handler for filename, size and headers
	file, handler, err := r.FormFile("myFile")
	if err != nil {
		fmt.Println("Error while handling file")
		fmt.Println(err)
		return
	}

	f, err := io.ReadAll(file)
	if err != nil {
		fmt.Printf("error while reading file: %v", err)
	}
	newFile, err := os.Create(handler.Filename)
	if err != nil {
		fmt.Printf("error while saving file: %v", err)
	}
	newFile.Write(f)

	defer file.Close()
	fmt.Printf("Uploaded File: %+v\n", handler.Filename)
	fmt.Printf("File Size: %+v\n", handler.Size)
	fmt.Printf("MIME Header: %+v\n", handler.Header)
	w.Write([]byte("Ok"))
}
