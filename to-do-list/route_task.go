package main

import (
	"net/http"
	"strconv"
	"to-do-list/data"
)

func profileDeleteTask(writer http.ResponseWriter, request *http.Request) {
	_, err := session(writer, request)
	if err != nil {
		http.Redirect(writer, request, "/login", 302)
		return
	}
	err = request.ParseForm()
	if err != nil {
		//danger method
	}
	id, err := strconv.ParseInt(request.PostFormValue("id"), 10, 10)
	err = data.DeleteTaskByID(int(id))

	http.Redirect(writer, request, "/profile-tasks", 302)
}

func profileAddTask(writer http.ResponseWriter, request *http.Request) {
	session, err := session(writer, request)
	if err != nil {
		http.Redirect(writer, request, "/login", 302)
		return
	}
	err = request.ParseForm()
	if err != nil {
		//danger method
	}
	task := data.Task{
		Title:       request.PostFormValue("title"),
		UserID:      session.User_ID,
		Deadline:    request.PostFormValue("deadline"),
		Description: request.PostFormValue("description"),
		IsImportant: isTrue(request.PostFormValue("isImportant")),
	}
	task.Create()
	http.Redirect(writer, request, "/profile-tasks", 302)
}
