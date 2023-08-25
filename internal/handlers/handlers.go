package handlers

import "net/http"

func SetupRouter() *http.ServeMux {
	router := http.NewServeMux()

	router.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello world!"))
	})

	return router
}
