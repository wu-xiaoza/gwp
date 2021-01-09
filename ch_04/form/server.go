package main

import (
	"fmt"
	"net/http")


func process(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Fprintln(w, r.Form)
	//PostForm只包含表单值，不包含URL里的值，并且只支持application/x-www-form-urlencoded编码
	// fmt.Fprintln(w, r.PostForm["post"])
	// fmt.Fprintln(w, r.PostForm["hello"])
}

func main() {
	server := http.Server{
Addr: "127.0.0.1:8080",
	}

	http.HandleFunc("/process", process)
	server.ListenAndServe()
}