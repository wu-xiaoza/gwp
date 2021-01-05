package main

import (
	"fmt"
	"net/http")


func writeExample(w http.ResponseWriter, r *http.Request) {
	str := `<html>
<head><title>Go Web Programming</title></head>
<body><h1>Hello World!</h1></body>
</html>`
	w.Write([]byte(str))
}

func writeHeaderExample(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(501)
	fmt.Fprintln(w, "No such service, try next door")
}

func headerExample(w http.ResponseWriter, r *http.Request) {
	//Set sets the header entries associated with key to the single element value. It replaces any existing values associated with key. The key is case insensitive; it is canonicalized by textproto.CanonicalMIMEHeaderKey. To use non-canonical keys, assign to the map directly.
	w.Header().Set("Location", "http://gopl.io")
	w.WriteHeader(302)
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/write", writeExample)
	http.HandleFunc("/redirect", headerExample)
	http.HandleFunc("/writeheader", writeHeaderExample)
	server.ListenAndServe()
}
