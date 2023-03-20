package main

import (
	"net/http"
	"to-do-list/data"
)

func main() {
	mux := http.NewServeMux()
	files := http.FileServer(http.Dir(data.Config.Static))
	mux.Handle("/static/", http.StripPrefix("/static/", files))

	privateFiles := http.FileServer(http.Dir(data.Config.Private))
	mux.Handle("/private/", http.StripPrefix("/private/", privateFiles))

	mux.HandleFunc("/", index)
	mux.HandleFunc("/sign-in", sign_in)
	mux.HandleFunc("/sign-up", sign_up)
	mux.HandleFunc("/login", login)
	mux.HandleFunc("/sign-up-account", signUpAccount)
	mux.HandleFunc("/logout", logout)
	mux.HandleFunc("/profile-tasks", profileTasks)
	mux.HandleFunc("/profile-add-task-page", profileAddTaskPage)
	mux.HandleFunc("/profile-add-task", profileAddTask)
	mux.HandleFunc("/profile-delete-task", profileDeleteTask)
	mux.HandleFunc("/edit-profile", settingsUpdateUserPage)
	mux.HandleFunc("/profile-update", settingsUpdateUser)
	mux.HandleFunc("/receipts-main-page", foodMainPage)
	mux.HandleFunc("/receipt-add-page", receiptAddPage)
	mux.HandleFunc("/receipt-add", receiptAdd)
	server := &http.Server{
		Addr:    data.Config.Address,
		Handler: mux,
	}
	server.ListenAndServe()
}

