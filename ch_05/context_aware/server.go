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
	server.ListenAndServe()
}

func process(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("tmpl.html")
	content := `I asked: <i>"What's up?"</i>`
	t.Execute(w, content)
}
/* 
$ curl -i localhost:8080/process
HTTP/1.1 200 OK
Date: Wed, 06 Jan 2021 09:29:53 GMT   
Content-Length: 544
Content-Type: text/html; charset=utf-8

<!DOCTYPE html>
<html>

<head>
    <meta http-equiv="Content-Type" content="text/html; charset=utf-8">
    <title>Go Web Programming</title>
</head>

<body>
    <div>I asked: &lt;i&gt;&#34;What&#39;s up?&#34;&lt;/i&gt;</div>
    <div><a href="/I%20asked:%20%3ci%3e%22What%27s%20up?%22%3c/i%3e">Path</a></div>    
    <div><a href="/?q=I%20asked%3a%20%3ci%3e%22What%27s%20up%3f%22%3c%2fi%3e">Query</a></div>
    <div><a onclick="f('I asked: \u003ci\u003e\u0022What\u0027s up?\u0022\u003c\/i\u003e')">Onclick</a></div>
</body>

</html>
	I asked: <i>"What's up?"</i>
	I asked: &lt;i&gt;&#34;What&#39;s up?&#34;&lt;/i&gt;
    I%20asked:  %20%3ci%3e%22What%27s%20up?  %22%3c/  i%3e
    I%20asked%3a%20%3ci%3e%22What%27s%20up%3f%22%3c%2fi%3e
    I asked: \u003ci\u003e\u0022What\u0027s up?\u0022\u003c\/i\u003e
*/
