package main

import "net/http"

func main() {

	mux := http.NewServeMux()

	registerRoutes(mux)
	http.ListenAndServe(":8080", mux)
}

type homeHandler struct{}

func (h *homeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("This is my home page"))
}

func registerRoutes(mux *http.ServeMux) {
	home := homeHandler{}
	mux.HandleFunc("/", home.ServeHTTP)
}
