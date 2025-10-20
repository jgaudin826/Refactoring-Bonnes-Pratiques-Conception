package services

import "net/http"

func Connect(w http.ResponseWriter, r *http.Request) {
	email := r.FormValue("email")
	SetCookie(email, w)
	if api.getUsers("data.json") {email} {
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}

func Disconnect(w http.ResponseWriter, r *http.Request) {
	DeleteCookie(w)
}
