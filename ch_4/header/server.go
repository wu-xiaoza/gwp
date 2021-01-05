package main

import (
	"fmt"
	"net/http")


func headers(w http.ResponseWriter, r *http.Request) {
	header := r.Header
	fmt.Fprintln(w, header)
	fmt.Fprintln(w)
	fmt.Fprintln(w)
	fmt.Fprintln(w)
	for h, c := range r.Header {
		fmt.Fprintf(w, "%v: %v\n", h, c)
	}
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/headers", headers)
	server.ListenAndServe()
}