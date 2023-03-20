package main

import (
	"errors"
	"io"
	"net/http"
	"os"
	"to-do-list/data"
)

func session(writer http.ResponseWriter, request *http.Request) (session data.Session, err error) {
	cookie, err := request.Cookie("_cookie")
	if err == nil {
		session = data.Session{UUID: cookie.Value}
		if ok, _ := session.Check(); !ok {
			err = errors.New("Invalid session")
		}
	}
	return
}

func isTrue(s string) bool {
	return s == "on"
}

func pasteFile(request *http.Request) (filename string, err error) {
	in, header, err := request.FormFile("avatar")
	if err != nil {
		return
	}
	defer in.Close()
	out, err := os.Create(data.Config.Private + "/avatar/" + header.Filename)

	if err != nil {
		return
	}
	defer out.Close()
	io.Copy(out, in)
	return out.Name(), nil
}

func pasteFileFood(request *http.Request) (filename string, err error) {
	in, header, err := request.FormFile("photo")
	if err != nil {
		return
	}
	defer in.Close()
	out, err := os.Create(data.Config.Private + "/food/" + header.Filename)

	if err != nil {
		return
	}
	defer out.Close()
	io.Copy(out, in)
	return out.Name(), nil
}
