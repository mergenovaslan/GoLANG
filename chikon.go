package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"html/template"
	"net/http"
)

type Main struct {
	Name    string
	Age     uint16
	Gender  string
	Balance int16
}

func (m Main) getAll() string {
	return fmt.Sprintf("Name is:%s,age is %d"+" Balance equal:%d", m.Name, m.Age, m.Balance)
}
func home(w http.ResponseWriter, r *http.Request) {
	chikon := Main{"Chikon", 20, "Female", 5000}
	templ, _ := template.ParseFiles("templates/home_page.html", "templates/second.html")
	templ.Execute(w, chikon)
}

func second(w http.ResponseWriter, r *http.Request) {
	sec := Main{"Diana", 20, "Female", 570}
	templ, _ := template.ParseFiles("templates/second.html")
	templ.Execute(w, sec)
}
func save(w http.ResponseWriter, r *http.Request) {
	ID := r.FormValue("ID")
	NAME := r.FormValue("NAME")
	SURNAME := r.FormValue("SURNAME")
	AGE := r.FormValue("AGE")

	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/golang")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	fmt.Println("Access")

	insert, err := db.Query(fmt.Sprintf("Insert Into 'users' ('ID','NAME','SURNAME','AGE') Values('%s','%s','%s','%s')", ID, NAME, SURNAME, AGE))
	if err != nil {
		panic(err)
	}
	defer insert.Close()

	http.RedirectHandler("/", 304)
}

func handleRequest() {
	http.HandleFunc("/", home)
	http.HandleFunc("/sec/", second)
	http.HandleFunc("/save", save)
	http.ListenAndServe(":5555", nil)
}
func main() {

	handleRequest()

}
