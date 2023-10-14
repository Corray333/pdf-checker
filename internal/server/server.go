package server

import "net/http"

func SetupRouter() *http.ServeMux {
	router := http.NewServeMux()

	router.HandleFunc("/test", LoadPDF)
	router.Handle("/", http.FileServer(http.Dir("../frontend/dist")))

	return router
}
