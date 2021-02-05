package main

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/google/uuid"
)

type user struct {
	FirstName string
	LastName  string
	UserName  string
}

var tpl *template.Template
var dbUsers = map[string]user{}     // userid -> user data
var dbSession = map[string]string{} // session id -> session user

func init() {
	tpl = template.Must(tpl.ParseFiles("index.html", "user.html"))
}

func main() {
	http.HandleFunc("/", home)
	http.HandleFunc("/user", userData)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func home(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("usersession")

	if err != nil {
		id := uuid.New()
		cookie = &http.Cookie{
			Name:     "usersession",
			Value:    id.String(),
			HttpOnly: true,
		}
		http.SetCookie(w, cookie)
	}

	var u user
	if currentUserID, ok := dbSession[cookie.Value]; ok {
		u = dbUsers[currentUserID]
	}

	if r.Method == "POST" {
		u.FirstName = r.FormValue("firstname")
		u.LastName = r.FormValue("lastname")
		u.UserName = r.FormValue("username")
		dbSession[cookie.Value] = u.UserName
		dbUsers[u.UserName] = u
	}
	tpl.ExecuteTemplate(w, "index.html", u)
}

func userData(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("usersession")
	if err != nil {
		fmt.Println("no usercookie found")
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	dbuser, ok := dbSession[cookie.Value]
	if !ok {
		fmt.Println("no user found")
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	fmt.Println("user found ", dbUsers[dbuser])
	tpl.ExecuteTemplate(w, "user.html", dbUsers[dbuser])
}
