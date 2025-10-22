package services

import (
	"net/http"
	"time"
)

// Gestion des Cookies
func SetCookie(value string, write http.ResponseWriter) {
	expiration := time.Now().Add(15 * 24 * time.Hour)
	cookie := http.Cookie{Name: "email", Value: value, Path: "/", Expires: expiration}
	http.SetCookie(write, &cookie)
}

func GetCookie(request *http.Request) string {
	var cookieUser *http.Cookie
	var errUser error

	cookieUser, errUser = request.Cookie("email")
	if errUser != nil {
		if errUser == http.ErrNoCookie {
			// No cookie = Not connected
			return ""
		}
	}
	return cookieUser.Value
}

func DeleteCookie(write http.ResponseWriter) {
	cookie := http.Cookie{Name: "email", Value: "", Path: "/", Expires: time.Unix(0, 0)}
	http.SetCookie(write, &cookie)
}
