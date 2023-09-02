package server

import "net/http"

func test(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello world!"))
}
