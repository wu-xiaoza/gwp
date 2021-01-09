package main

import (
	"fmt"
	"net/http")


func process(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, r.FormValue("hello"))
}

func main() {
	server := http.Server{
Addr: "127.0.0.1:8080",
	}

	http.HandleFunc("/process", process)
	server.ListenAndServe()
}