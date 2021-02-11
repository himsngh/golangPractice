package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
)

type person struct {
	FName string
	LName string
	Age   int
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("index.html"))
}

func main() {
	http.HandleFunc("/", home)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.HandleFunc("/marshal", marshal)
	http.HandleFunc("/unmarshal", unmarshal)
	http.HandleFunc("/encode", encode)
	http.ListenAndServe(":8080", nil)
}

func home(w http.ResponseWriter, r *http.Request) {
	tpl.Execute(w, nil)
}

func marshal(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	p := person{
		FName: "Mount",
		LName: "Everton",
		Age:   22,
	}

	jsonMarshal, err := json.Marshal(p)
	if err != nil {
		fmt.Println(err)
		return
	}
	if err != nil {
		fmt.Println(err)
		return
	}
	w.Write(jsonMarshal)
}

func unmarshal(w http.ResponseWriter, r *http.Request) {
	p := person{
		FName: "Mount",
		LName: "Everton",
		Age:   22,
	}
	jsonMarshal, err := json.Marshal(p)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(jsonMarshal)
	if err != nil {
		fmt.Println(err)
		return
	}
	var unmarshalP person
	err = json.Unmarshal(jsonMarshal, &unmarshalP)
	tpl.Execute(w, unmarshalP)
}

func encode(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	p := person{
		FName: "Rocky",
		LName: "Mountain",
		Age:   22,
	}
	err := json.NewEncoder(w).Encode(&p)
	if err != nil {
		fmt.Println(err)
	}
}
