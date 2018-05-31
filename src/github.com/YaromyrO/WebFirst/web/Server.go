package main

import (
	"net/http"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir("src/github.com/YaromyrO/WebFirst/html")))
	http.ListenAndServe(":8000", nil)
}

