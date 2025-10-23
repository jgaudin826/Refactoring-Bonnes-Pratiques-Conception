package services

import (
	"net/http"
	"refactoring/api"
)

func Connect(write http.ResponseWriter, request *http.Request) {
	userEmail := request.FormValue("email")
	SetCookie(userEmail, write)
	userExists := false
	userList := api.GetUsers()
	for _, user := range userList {
		if user.Email == userEmail {
			userExists = true
			break
		}
	}
	if !userExists {
		newUser := api.User{
			ID:    len(userList) + 1,
			Email: userEmail,
			Role:  "user",
		}
		api.AddUser(newUser)
	}
	http.Redirect(write, request, "/", http.StatusSeeOther)
}

func Disconnect(write http.ResponseWriter, request *http.Request) {
	DeleteCookie(write)
	http.Redirect(write, request, "/", http.StatusSeeOther)
}
