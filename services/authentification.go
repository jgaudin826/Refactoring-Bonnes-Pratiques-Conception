package services

import (
	"net/http"
	"refactoring/api"
)

func Connect(w http.ResponseWriter, r *http.Request) {
	email := r.FormValue("email")
	SetCookie(email, w)
	emailExists := false
	userList := api.GetUsers("data.json")
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
		api.AddUser("data.json", newUser)
	}
}

func Disconnect(w http.ResponseWriter, r *http.Request) {
	DeleteCookie(w)
}
