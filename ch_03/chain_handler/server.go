package main

import (
	"fmt"
	"net/http"
)

//HelloHandler .
type HelloHandler struct{}

func (h HelloHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello!")
}

func log(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("Handler called - %T\n", h)
		h.ServeHTTP(w, r)
	})
}

func protect(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("Protect protect - %T\n", h)
		h.ServeHTTP(w, r)
	})
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	hello := HelloHandler{}
	//HandleFunc registers the handler function for the given pattern in the DefaultServeMux. The documentation for ServeMux explains how patterns are matched.
	http.Handle("/hello", protect(log(hello)))
	server.ListenAndServe()
	/* http.Handle()
	http.Handler
	http.HandleFunc()
	http.HandlerFunc */
}

