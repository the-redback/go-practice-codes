package main

import (
	"net/http"
	"io"
)

func hello(res http.ResponseWriter, req *http.Request) {
	res.Header().Set(
		"Content-Type",
		"Text/html",
	)
	io.WriteString(
		res,
		`<doctype html>
	<html>
		<head>
		<title> Hello World</title>
		</head>
		<body>
			Hello World
		</body>

	</html>`,
	)
}


//Can't see any output of this programm
//Have no idea why!!! -_-

func main() {
	http.HandleFunc("/hello",hello)
	http.ListenAndServe(":9000",nil)

	//File Server

	http.Handle(
		"/assets/",
		http.StripPrefix(
			"/assets/",
			http.FileServer(http.Dir("assets")),
		),
	)
}