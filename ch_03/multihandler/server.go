package main

import (
	"fmt"
	"net/http"
)

//HelloHandler .
type HelloHandler struct{}

func (h *HelloHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello!")
}

//WorldHandler .
type WorldHandler struct{}

func (h *WorldHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Worlld!")
}

func main() {
	hello := HelloHandler{}
	world := WorldHandler{}
	server := http.Server{
		Addr:    "127.0.0.1:8080",
	}
	//Handle registers the handler for the given pattern in the DefaultServeMux. The documentation for ServeMux explains how patterns are matched.
	http.Handle("/hello", &hello)
	http.Handle("/world", &world)
	server.ListenAndServe()
}