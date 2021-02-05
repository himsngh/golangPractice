package main

import (
	"io"
	"net/http"
)

func main() {
	// http.FileServer(http.Dir(".")) -> it is serving everything in the curent directory
	//  That means even the file main.go is served
	// http.Dir("./..")	server everything to the root folder of this file
	// http.Handle("/", http.FileServer(http.Dir("/Users/himanshsingh/Desktop/golang/Servers")))
	http.Handle("/", http.FileServer(http.Dir("./..")))
	http.HandleFunc("/dog", dog)

	http.ListenAndServe(":8080", nil)

}

func dog(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "text/html; character-set=utf-8")
	io.WriteString(w, ` <img src="dog-img.jpg" >`)
}
