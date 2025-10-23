package services

import (
	"net/http"
	"refactoring/api"
)

func Connect(write http.ResponseWriter, request *http.Request) {
	email := request.FormValue("email")
	SetCookie(email, write)
	emailExists := false
	userList := api.GetUsers()
	for _, user := range userList {
		if user.Email == email {
			emailExists = true
			break
		}
	}
	if !emailExists {
		newUser := api.User{
			ID:    len(userList) + 1,
			Email: email,
			Role:  "user",
		}
		api.AddUser(newUser)
	}
	http.Redirect(write, request, "/", http.StatusSeeOther)
	return
}

func Disconnect(write http.ResponseWriter, request *http.Request) {
	DeleteCookie(write)
	http.Redirect(write, request, "/", http.StatusSeeOther)
	return
}
