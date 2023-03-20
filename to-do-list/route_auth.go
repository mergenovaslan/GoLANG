package main

import (
	"html/template"
	"net/http"
	"to-do-list/data"
)

func signUpAccount(writer http.ResponseWriter, request *http.Request) {
	err := request.ParseForm()
	if err != nil {
		//danger method
	}
	if request.PostFormValue("password") == request.PostFormValue("repassword") {
		user := data.User{
			Username: request.PostFormValue("username"),
			Name:     request.PostFormValue("name"),
			Email:    request.PostFormValue("email"),
			Password: request.PostFormValue("password"),
		}
		if err := user.Create(); err != nil {
			//danger method
		}
		http.Redirect(writer, request, "/sign-in", 302)
	} else {
		//danger method
	}
}

func sign_in(writer http.ResponseWriter, request *http.Request) {
	t, _ := template.ParseFiles("templates/sign-in.html")
	t.Execute(writer, nil)
}

func sign_up(writer http.ResponseWriter, request *http.Request) {
	t, _ := template.ParseFiles("templates/sign-up.html")
	t.Execute(writer, nil)
}

func login(writer http.ResponseWriter, request *http.Request) {
	err := request.ParseForm()
	user, err := data.UserByEmailOrUsername(request.PostFormValue("username-or-email"))
	if err != nil {
		//danger method
	}
	if user.Password == data.Encrypt(request.PostFormValue("password")) {
		session, err := user.CreateSession()
		if err != nil {
			//danger method
		}
		cookie := http.Cookie{
			Name:     "_cookie",
			Value:    session.UUID,
			HttpOnly: true,
		}
		http.SetCookie(writer, &cookie)
		http.Redirect(writer, request, "/", 302)
	} else {
		http.Redirect(writer, request, "/sign-in", 302)
	}

}
func logout(writer http.ResponseWriter, request *http.Request) {
	cookie, err := request.Cookie("_cookie")
	if err != http.ErrNoCookie {
		//warning method
		session := data.Session{UUID: cookie.Value}
		session.DeleteByUUID()
	}
	http.Redirect(writer, request, "/", 302)
}
