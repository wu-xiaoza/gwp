package main

import (
	"html/template"
	"math/rand"
	"net/http"
	"time"
)


func main() {
	server := http.Server{
		Addr: "localhost:8080",
	}
	http.HandleFunc("/process", process)
	server.ListenAndServe()
}

func process(w http.ResponseWriter, r *http.Request) {
	rand.Seed(time.Now().Unix())
	var t *template.Template
	if rand.Intn(10) > 5 {
		t, _ = template.ParseFiles("layout.html", "red_hello.html")
		} else {
		t, _ = template.ParseFiles("layout.html", "blue_hello.html")
	}
	t.ExecuteTemplate(w, "layout", "")
}