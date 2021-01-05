package main

import (
	"fmt"
	"net/http")


func setCookie(w http.ResponseWriter, r *http.Request) {
	c1 := http.Cookie{
		Name: "first_cookie",
		Value: "Co Web Programming",
		HttpOnly: true,
	}
	c2 := http.Cookie{
		Name: "second_cookie",
		Value: "Manning Publications Co",
		HttpOnly: true,
	}
	http.SetCookie(w, &c1)
	http.SetCookie(w, &c2)
}

func getCookie(w http.ResponseWriter, r *http.Request) {
	c1, err := r.Cookie("first_cookie")
	if err != nil {
		fmt.Fprintln(w, "Cannot get the first cookie")
	}
	c2, err := r.Cookie("second_cookie")
	if err != nil {
		fmt.Fprintln(w, "Cannot get the second cookie")
	}
	c3, err := r.Cookie("third_cookie")
	if err != nil {
		fmt.Fprintln(w, "Cannot get the third cookie")
	}
	cs := r.Cookies()
	fmt.Fprintln(w, c1)
	fmt.Fprintln(w, c2)
	fmt.Fprintln(w, c3)
	fmt.Fprintln(w, cs)

	for _, v := range cs {
		fmt.Fprintln(w, v)
	}

}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/set_cookie", setCookie)
	http.HandleFunc("/get_cookie", getCookie)
	server.ListenAndServe()
}