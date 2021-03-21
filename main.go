package main

import (
	"log"
	"net/http"
	"strconv"
	"text/template"

	"github.com/gorilla/mux"
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
	if r.Method == "POST" {
		vars := mux.Vars(r)
		id, _ := strconv.Atoi(vars["id"])
		UList.List[id-1].Access = true
		http.Redirect(w, r, "/info", http.StatusMovedPermanently)
	}
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/info", infoHandler)
	router.HandleFunc("/access/{id}", accHandler).Methods("POST")
	router.HandleFunc("/", indexHandler)
	http.Handle("/", router)
	log.Fatal(http.ListenAndServe(":8081", nil))
}
