package main

import (
	"fmt"
	"os"
	"strings"
	"text/template"
)

var fm = template.FuncMap{

	"func1": strings.ToUpper,
	"func2": myFunc,
}

func myFunc(str string) string {
	str = str + "\t\tHello , How are you ?"
	return str
}

var tpl *template.Template

func init() {
	// if we don't provide any names in the New("") then we have to call ExecuteTemplate with the name of the file we want to execute
	//tpl = template.Must(template.New("").Funcs(fm).ParseFiles("go.html"))
	tpl = template.Must(template.New("go.html").Funcs(fm).ParseFiles("go.html"))
}

func main() {
	data := []string{
		"himanshu",
		"himansh ",
		"alex    ",
		"tj      ",
		"ranjan  ",
	}

	//err := tpl.ExecuteTemplate(os.Stdout, "go.html", data)
	err := tpl.Execute(os.Stdout, data)
	if err != nil {
		fmt.Println(err)
	}
}
