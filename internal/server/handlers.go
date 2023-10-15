package server

import (
	"bytes"
	"encoding/json"
	"io"
	"log/slog"
	"net/http"
	"os"
	"pdfchecker/internal/validation"
)

func LoadPDFWrapper(logger *slog.Logger) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		r.ParseMultipartForm(10 << 20)

		file, handler, err := r.FormFile("myFile")
		if err != nil {
			logger.Error("error while handling file:" + err.Error())
			return
		}

		f, err := io.ReadAll(file)
		if err != nil {
			logger.Error("error while reading file:" + err.Error())
			return
		}
		newFile, err := os.Create("../pdf/" + handler.Filename)
		if err != nil {
			logger.Error("error while saving file:" + err.Error())
			return
		}
		newFile.Write(f)

		defer file.Close()

		params, _ := json.Marshal(struct {
			FileName string `json:"fileName"`
		}{FileName: "../pdf/" + handler.Filename})
		resp, err := http.Post("https://httpbin.org/post", "application/json", bytes.NewBuffer(params))
		if err != nil {
			logger.Error("error while connection to pdf server:" + err.Error())
		}
		defer resp.Body.Close()
		respText, err := io.ReadAll(resp.Body)
		if err != nil {
			logger.Error("error while reading pdf errors:" + err.Error())
		}
		errors, err := json.Marshal(validation.ErrorDetection(string(respText)))
		if err != nil {
			logger.Error("error while marshalling errors:" + err.Error())
		}
		w.Write(errors)
	}
}
