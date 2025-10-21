package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"refactoring/api"
	"refactoring/services"
)

var port = ":8080"

func home(w http.ResponseWriter, r *http.Request) {

	tmpl, err := template.ParseFiles("./templates/home.html") // Read the home page
	if err != nil {
		log.Printf("\033[31mError parsing template: %v\033[0m", err)
		http.Error(w, "Internal error, template not found.", http.StatusInternalServerError)
		return
	}
	var user api.User
	email := services.GetCookie(r)
	if email == "" {

	} else {
		userList := api.GetUsers("data/data.json")
		for _, users := range userList {
			if users.Email == email {
				user = users
				break
			}
		}
	}

	homePage := struct {
		User     api.User
		Services []api.Service
		Bookings []api.Booking
	}{
		User:     user,
		Services: api.GetServices("data/data.json"),
		Bookings: services.GetBookingsByEmail("data/data.json", email),
	}

	err = tmpl.Execute(w, homePage)
	if err != nil {
		log.Printf("\033[31mError executing template: %v\033[0m", err)
		http.Error(w, "Internal error", http.StatusInternalServerError)
		return
	}
}

func main() {
	FileServer := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", FileServer))

	http.HandleFunc("/", home)

	// FORMS
	http.HandleFunc("/Connect", services.Connect)
	http.HandleFunc("/Disconnect", services.Disconnect)
	http.HandleFunc("/AddService", services.AddService)
	http.HandleFunc("/AddSlot", services.AddSlot)
	http.HandleFunc("/CancelBooking", services.CancelBooking)
	http.HandleFunc("/BookingSlot", services.BookingSlot)

	fmt.Println("Server Start at:")
	fmt.Println("http://localhost" + port)

	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatal(err)
	}
}
