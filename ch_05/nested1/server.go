package main

import (
	"net/http"	
	"html/template"
)


func main() {
	server := http.Server{
		Addr: "localhost:8080",
	}
	http.HandleFunc("/process", process)
	server.ListenAndServe()
}

func process(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("layout.html")
	t.ExecuteTemplate(w, "layout", "")
}