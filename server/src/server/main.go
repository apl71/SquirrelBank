package main

import (
	"net/http"
)

var www_path string = "../../../www/"

func main() {
	http.Handle("/", http.FileServer(http.Dir(www_path)))
	http.ListenAndServe(":80", nil)
}