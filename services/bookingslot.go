package services

import (
	"fmt"
	"net/http"
	"refactoring/api"
	"strconv"
)

// Book slot from user submitted form
// Check if a user is connected and if slot and service exists
// Calls AddBooking and RemoveSlotFromService
func BookSlot(write http.ResponseWriter, request *http.Request) {
	userEmail := GetCookie(request)
	serviceIDStr := request.FormValue("servicesId")
	selectedSlot := request.FormValue("slot")

	if userEmail == "" {
		fmt.Println("Aucun email trouvé dans les cookies.")
		http.Redirect(write, request, "/", http.StatusSeeOther)
	}

	if serviceIDStr == "" {
		fmt.Println("Aucun ID de service fourni.")
		http.Redirect(write, request, "/", http.StatusSeeOther)
	}

	if selectedSlot == "" {
		fmt.Println("Aucun créneau sélectionné.")
		http.Redirect(write, request, "/", http.StatusSeeOther)
	}

	bookings := api.GetBookings()
	services := api.GetServices()

	serviceID, convertError := strconv.Atoi(serviceIDStr)
	if convertError != nil {
		fmt.Println("L'ID du service est invalide.")
		http.Redirect(write, request, "/", http.StatusSeeOther)
	}

	var selectedService *api.Service
	for index, service := range services {
		if service.ID == serviceID {
			selectedService = &services[index]
			break
		}
	}

	if selectedService == nil {
		fmt.Println("Service introuvable.")
		http.Redirect(write, request, "/", http.StatusSeeOther)
	}

	slotAvailable := false
	for _, slot := range selectedService.Slots {
		if slot == selectedSlot {
			slotAvailable = true
			break
		}
	}
	if !slotAvailable {
		fmt.Println("Ce créneau n'existe pas pour ce service.")
		http.Redirect(write, request, "/", http.StatusSeeOther)
	}

	for _, booking := range bookings {
		if booking.Email == userEmail && booking.Service == serviceID && booking.Slot == selectedSlot {
			fmt.Println("Vous avez déjà réservé ce créneau.")
			http.Redirect(write, request, "/", http.StatusSeeOther)
		}
	}

	for _, booking := range bookings {
		if booking.Service == serviceID && booking.Slot == selectedSlot {
			fmt.Println("Ce créneau est déjà complet.")
			http.Redirect(write, request, "/", http.StatusSeeOther)
		}
	}

	newBooking := api.Booking{
		ID:      len(bookings) + 1,
		Email:   userEmail,
		Service: serviceID,
		Slot:    selectedSlot,
	}

	api.AddBooking(newBooking)
	api.RemoveSlotFromService(newBooking.Service, selectedSlot)
	http.Redirect(write, request, "/", http.StatusSeeOther)
}
