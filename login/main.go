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
	http.HandleFunc("/logout", logOut)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func home(w http.ResponseWriter, r *http.Request) {
	if loggedIN(r) {
		fmt.Println("Logged In")
		http.Redirect(w, r, "/user", http.StatusSeeOther)
		return
	}

	var u user
	if r.Method == "POST" {
		id := uuid.New()
		cookie := &http.Cookie{
			Name:     "usersession",
			Value:    id.String(),
			HttpOnly: true,
		}
		http.SetCookie(w, cookie)
		u.FirstName = r.FormValue("firstname")
		u.LastName = r.FormValue("lastname")
		u.UserName = r.FormValue("username")
		dbSession[cookie.Value] = u.UserName
		dbUsers[u.UserName] = u
	}
	tpl.ExecuteTemplate(w, "index.html", u)
}

func userData(w http.ResponseWriter, r *http.Request) {
	if !loggedIN(r) {
		fmt.Println("Not logged In ")
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	cookie, _ := r.Cookie("usersession")
	dbuser, _ := dbSession[cookie.Value]
	fmt.Println("user found ", dbUsers[dbuser])
	tpl.ExecuteTemplate(w, "user.html", dbUsers[dbuser])
}

func loggedIN(r *http.Request) bool {
	c, err := r.Cookie("usersession")
	if err != nil {
		fmt.Println("No user session started")
		return false
	}
	un, ok := dbSession[c.Value]
	if !ok {
		fmt.Println("No user exist")
		return false
	}
	_, ok = dbUsers[un]
	return ok
}

func logOut(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("usersession")
	if err != nil {
		fmt.Println("No user session")
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	un, ok := dbSession[c.Value]
	if !ok {
		fmt.Println("No user")
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	delete(dbUsers, un)
	delete(dbSession, c.Value)
	fmt.Println("logged Out")
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
