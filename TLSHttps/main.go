package main

import (
	"fmt"
	"net/http"
)

// go run usr/local/go/src/crypto/tls/generate_cert.go
func main() {
	http.HandleFunc("/", foo)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServeTLS(":10443", "cert.pem", "key.pem", nil)
}

func foo(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Https Server up and running")
}
