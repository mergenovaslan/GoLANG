package main

import (
	"html/template"
	"net/http"
	"to-do-list/data"
)

type UserSettings struct {
	Username string
	Name     string
	Email    string
	Avatar   string
}

func profileSettings(writer http.ResponseWriter, request *http.Request) {
	session, err := session(writer, request)
	if err != nil {
		http.Redirect(writer, request, "/login", 302)
		return
	}
	user, err := data.UserByID(session.User_ID)
	t, err := template.ParseFiles(
		"templates/base.html", "templates/private.navbar.html",
		"templates/profile-base", "templates/profiles-settings.html",
	)
	if err != nil {
		//danger method
		return
	}
	t.ExecuteTemplate(writer, "base", user)
}

func settingsUpdateUserPage(writer http.ResponseWriter, request *http.Request) {
	session, err := session(writer, request)
	if err != nil {
		http.Redirect(writer, request, "/login", 302)
		return
	}
	user, err := data.UserByID(session.User_ID)
	if err != nil {
		//danger method
		return
	}
	t, err := template.ParseFiles(
		"templates/base.html", "templates/private.navbar.html",
		"templates/profile-base.html", "templates/profile-settings-update.html",
	)
	if err != nil {
		//danger method
		return
	}
	t.ExecuteTemplate(writer, "base", user)
	return
}
func settingsUpdateUser(writer http.ResponseWriter, request *http.Request) {
	request.ParseMultipartForm(10 << 20)
	session, err := session(writer, request)
	user, err := data.UserByID(session.User_ID)
	if err != nil {
		// danger method
		return
	}
	if data.Encrypt(request.PostFormValue("password")) == user.Password {
		user.Name = request.PostFormValue("name")
		user.Email = request.PostFormValue("email")
		user.Username = request.PostFormValue("username")
		filename, err := pasteFile(request)
		if err == nil {
			user.Avatar = filename
		}
		err = user.Update()
		if err != nil {
			//danger method
			return
		}
	}
	http.Redirect(writer, request, "/edit-profile", 302)
}
