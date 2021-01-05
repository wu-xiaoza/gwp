package main

import (
	"fmt"
	"net/http")


func index(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(rw, "hello world")
}

func hello(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(rw, "hello")
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", index)
	mux.HandleFunc("/hello/", hello)  /* /hello  访问/hello/XXX会返回index处理器处理结果     /hello/  访问/hello/XXX会返回hello处理器处理结果 */
	server := &http.Server{
		Addr: "127.0.0.1:8080",
		Handler: mux,
	}
	server.ListenAndServe()
}