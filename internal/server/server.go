package server

import (
	"log/slog"
	"net/http"
)

func SetupRouter(app *slog.Logger) *http.ServeMux {
	router := http.NewServeMux()

	router.HandleFunc("/loadpdf", LoadPDFWrapper(app))
	router.Handle("/", http.FileServer(http.Dir("../frontend/dist")))

	return router
}
