package main

import (
	"fmt"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Home Page!!!")
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
