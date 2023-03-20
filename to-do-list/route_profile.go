package main

import (
	"html/template"
	"net/http"
	"to-do-list/data"
)

func profileTasks(writer http.ResponseWriter, request *http.Request) {
	session, err := session(writer, request)
	if err != nil {
		http.Redirect(writer, request, "/login", 302)
		return
	}
	tasks, err := data.UserTasksByUserID(session.User_ID)
	t, err := template.ParseFiles(
		"templates/base.html", "templates/private.navbar.html",
		"templates/profile-base.html", "templates/profile-tasks.html",
	)
	t.ExecuteTemplate(writer, "base", tasks)
}

func profileAddTaskPage(writer http.ResponseWriter, request *http.Request) {
	_, err := session(writer, request)
	if err != nil {
		http.Redirect(writer, request, "/login", 302)
		return
	}
	t, err := template.ParseFiles(
		"templates/base.html", "templates/private.navbar.html", "templates/profile-base.html",
		"templates/profile-task-add.html",
	)
	t.ExecuteTemplate(writer, "base", nil)
}
