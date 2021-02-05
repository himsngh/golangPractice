package main

import (
	"io"
	"net/http"
	"os"
)

func main() {

	http.HandleFunc("/", dog)
	http.HandleFunc("/dog-img.jpg", dogPic)

	http.ListenAndServe(":8080", nil)
}

func dog(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "text/html; characterset=utf-8")
	io.WriteString(w, `  
	<img src = "/dog-img.jpg">
	`)
}

func dogPic(w http.ResponseWriter, r *http.Request) {
	// One way to do it
	file, err := os.Open("dog-img.jpg")

	if err != nil {
		io.WriteString(os.Stdout, ` "File not found ", err `)
		http.Error(w, "File not found ", 404)
	}
	defer file.Close()
	io.Copy(w, file)

	// Another way using serveFile()
	// http.ServeFile(w, r, "dog-img.jpg")
}
