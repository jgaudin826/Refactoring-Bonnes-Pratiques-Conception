package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
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

	email := api.GetCookie(r)
	if email == "" {
		email = "non connecté"
	} else {
		
	}

	homePage := struct {
		User services.User
		Services []services.Service
		Bookings []services.Booking

	}{
		//User: 
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
	http.HandleFunc("/connect", services.Connect)
	http.HandleFunc("/disconnect", services.Disconnect)
	http.HandleFunc("/AddService", services.AddService)
	http.HandleFunc("/AddSlot", services.AddSlot)
	http.HandleFunc("/CancelBooking", services.CancelBooking)

	fmt.Println("Server Start at:")
	fmt.Println("http://localhost" + port)

	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatal(err)
	}
}
