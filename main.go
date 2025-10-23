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

func home(write http.ResponseWriter, request *http.Request) {
	homeTemplate, parseError := template.ParseFiles("./templates/home.html") // Lecture du template de la page d'accueil
	if parseError != nil {
		log.Printf("\033[31mErreur lors du parsing du template: %v\033[0m", parseError)
		http.Error(write, "Erreur interne : template introuvable.", http.StatusInternalServerError)
		return
	}

	var currentUser api.User
	var userBookings []api.Booking

	userEmail := services.GetCookie(request)
	if userEmail != "" {
		users := api.GetUsers()
		for _, user := range users {
			if user.Email == userEmail {
				currentUser = user
				break
			}
		}
	}

	if currentUser.Role == "admin" {
		userBookings = api.GetBookings()
	} else {
		userBookings = api.GetBookingsByEmail(userEmail)
	}

	pageData := struct {
		User     api.User
		Services []api.Service
		Bookings []api.Booking
	}{
		User:     currentUser,
		Services: api.GetServices(),
		Bookings: userBookings,
	}

	execError := homeTemplate.Execute(write, pageData)
	if execError != nil {
		log.Printf("\033[31mErreur lors de l'exécution du template: %v\033[0m", execError)
		http.Error(write, "Erreur interne", http.StatusInternalServerError)
		return
	}
}

func main() {
	fileServer := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fileServer))

	http.HandleFunc("/", home)

	// ROUTES DES FORMULAIRES
	http.HandleFunc("/Connect", services.Connect)
	http.HandleFunc("/Disconnect", services.Disconnect)
	http.HandleFunc("/CreateService", services.CreateService)
	http.HandleFunc("/CreateSlot", services.CreateSlot)
	http.HandleFunc("/CancelBooking", services.CancelBooking)
	http.HandleFunc("/BookingSlot", services.BookSlot)

	fmt.Println("Server started at:")
	fmt.Println("👉 http://localhost" + port)

	startError := http.ListenAndServe(port, nil)
	if startError != nil {
		log.Fatal(startError)
	}
}
