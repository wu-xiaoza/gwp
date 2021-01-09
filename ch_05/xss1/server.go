package main

import (
	"html/template"
	"net/http"
)

func main() {
	server := http.Server{
		Addr: "localhost:8080",
	}
	http.HandleFunc("/process", process)
	http.HandleFunc("/form", form)
	server.ListenAndServe()
}

func process(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("X-XSS-Protection", "0")
	t, _ := template.ParseFiles("tmpl.html") 
	t.Execute(w, template.HTML(r.FormValue("comment")))
}

func form(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("form.html")
	t.Execute(w, nil)
}