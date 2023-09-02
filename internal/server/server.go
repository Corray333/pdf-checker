package server

import "net/http"

func SetupRouter() *http.ServeMux {
	router := http.NewServeMux()

	router.HandleFunc("/test", test)

	return router
}
