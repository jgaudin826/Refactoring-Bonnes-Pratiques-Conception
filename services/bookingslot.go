package services

import (
	"fmt"
	"net/http"
	"refactoring/api"
	"strconv"
)

func BookSlot(write http.ResponseWriter, request *http.Request) {
	email := GetCookie(request)
	ID := request.FormValue("servicesId")
	slot := request.FormValue("slot")
	if email == "" {
		fmt.Println("il n'y a pas d'email")
		http.Redirect(write, request, "/", http.StatusSeeOther)
		return
	} else if ID == "" {
		fmt.Println("il n'y a pas d'ID de service")
		http.Redirect(write, request, "/", http.StatusSeeOther)
		return
	} else if slot == "" {
		fmt.Println("il n'y a pas de créneau")
		http.Redirect(write, request, "/", http.StatusSeeOther)
		return
	}
	dataBookings := api.GetBookings()
	dataServices := api.GetServices()
	serviceID, err := strconv.Atoi(ID)
	if err != nil {
		fmt.Println("L'id n'est pas valide !")
		http.Redirect(write, request, "/", http.StatusSeeOther)
		return
	}
	var service *api.Service
	for index, services := range dataServices {
		if services.ID == serviceID {
			service = &dataServices[index]
			break
		}
	}
	if service == nil {
		fmt.Println("Service introuvable.")
		http.Redirect(write, request, "/", http.StatusSeeOther)
		return
	}

	slotExists := false
	for _, slots := range service.Slots {
		if slots == slot {
			slotExists = true
			break
		}
	}
	if !slotExists {
		fmt.Println("Ce créneau n'existe pas pour ce service.")
		http.Redirect(write, request, "/", http.StatusSeeOther)
		return
	}

	for _, bookings := range dataBookings {
		if bookings.Email == email && bookings.Service == serviceID && bookings.Slot == slot {
			fmt.Println("Vous avez déjà réservé ce créneau.")
			http.Redirect(write, request, "/", http.StatusSeeOther)
			return
		}
	}

	for _, bookings := range dataBookings {
		if bookings.Service == serviceID && bookings.Slot == slot {
			fmt.Println("Ce créneau est déjà complet.")
			http.Redirect(write, request, "/", http.StatusSeeOther)
			return
		}
	}
	newBooking := api.Booking{
		ID:      len(dataBookings) + 1,
		Email:   email,
		Service: serviceID,
		Slot:    slot,
	}
	api.AddBooking(newBooking)
	api.RemoveSlotFromService(newBooking.Service, slot)
	http.Redirect(write, request, "/", http.StatusSeeOther)
}

func GetBookingsByEmail(email string) []api.Booking {
	dataBookings := api.GetBookings()
	var userBookings []api.Booking

	for _, booking := range dataBookings {
		if booking.Email == email {
			userBookings = append(userBookings, booking)
		}
	}
	return userBookings
}
