package main

import (
	"fmt"
	"html/template"
	"net/http"
)

var t *template.Template

func init() {
	t = template.Must(template.ParseFiles("templates/header.html", "templates/index.html", "templates/footer.html"))
}

type person struct {
	FirstName string
	LastName  string
	Subscribe bool
}

func main() {
	http.HandleFunc("/index", index)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.Handle("/", http.FileServer(http.Dir(".")))
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	fName := r.FormValue("firstName")
	lName := r.FormValue("lastName")
	sub := r.FormValue("subscribe") == "on"

	err := t.ExecuteTemplate(w, "index.html", person{fName, lName, sub})
	if err != nil {
		fmt.Println("Error : ", err)
	}
}
