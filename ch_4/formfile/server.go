package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)


func process(w http.ResponseWriter, r *http.Request) {
	file, _ , err:= r.FormFile("uploaded")
	if err == nil {
		//ReadAll reads from r until an error or EOF and returns the data it read. A successful call returns err == nil, not err == EOF. Because ReadAll is defined to read from src until EOF, it does not treat an EOF from Read as an error to be reported.
		data, err := ioutil.ReadAll(file)
		if err == nil {
			fmt.Fprintln(w, string(data))
		}
	}
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/process", process)
	server.ListenAndServe()
}