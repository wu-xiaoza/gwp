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
	t, _ := template.ParseFiles("tmpl.html")
	//Seed uses the provided seed value to initialize the default Source to a deterministic state. If Seed is not called, the generator behaves as if seeded by Seed(1). Seed values that have the same remainder when divided by 2³¹-1 generate the same pseudo-random sequence. Seed, unlike the Rand.Seed method, is safe for concurrent use.
	rand.Seed(time.Now().Unix())//保证随机数不重复
	//Intn returns, as an int, a non-negative pseudo-random number in [0,n) from the default Source. It panics if n <= 0.
	t.Execute(w, rand.Intn(10) > 5)
}

