package main

import (
	"fmt"
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
	//fmt.Fprintf(w, chikon.getAll())
	templ, _ := template.ParseFiles("templates/home_page.html")
	templ.Execute(w, chikon)
}

func second(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Second Page")
}

func handleRequest() {
	http.HandleFunc("/", home)
	http.HandleFunc("/sec/", second)
	http.ListenAndServe(":5555", nil)
}
func main() {
	handleRequest()
}
