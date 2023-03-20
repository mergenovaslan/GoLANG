package main

import (
	"html/template"
	"net/http"
)

func index(writer http.ResponseWriter, request *http.Request) {
	_, err := session(writer, request)
	if err != nil {
		t, _ := template.ParseFiles("templates/base.html", "templates/public.navbar.html", "templates/index.html")
		t.ExecuteTemplate(writer, "base", nil)
	} else {
		t, _ := template.ParseFiles("templates/base.html", "templates/private.navbar.html", "templates/index.html")
		t.ExecuteTemplate(writer, "base", nil)
	}
}
