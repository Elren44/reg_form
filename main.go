package main

import (
	"log"
	"net/http"
	"strconv"
	"text/template"
)

type User struct {
	ID       int
	Name     string
	Password string
	Access   bool
}

type UserList struct {
	List []User
}

var UList UserList

func (ul *UserList) newUser(id int, n string, p string) {
	x := User{ID: id, Name: n, Password: p}
	ul.List = append(ul.List, x)
}

var indexHTML = template.Must(template.ParseFiles("index.html"))

func indexHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		name := r.FormValue("login")
		pass := r.FormValue("password")
		id := len(UList.List) + 1
		UList.newUser(id, name, pass)
	}
	indexHTML.Execute(w, UList)
}

var infoHTML = template.Must(template.ParseFiles("info.html"))

func infoHandler(w http.ResponseWriter, r *http.Request) {
	infoHTML.Execute(w, UList)
}

func accHandler(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(r.URL.Path[len("/access/"):])
	for _, val := range UList.List {
		if val.ID == id {
			UList.List[id-1].Access = true
		}
	}
	http.Redirect(w, r, "/info", http.StatusPermanentRedirect)
}

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/info/", infoHandler)
	http.HandleFunc("/access/", accHandler)
	log.Fatal(http.ListenAndServe(":8081", nil))
}
