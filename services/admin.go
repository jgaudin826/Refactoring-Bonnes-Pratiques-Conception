package services

import (
	"fmt"
	"net/http"
	"refactoring/api"
	"strconv"
)

func AddService(w http.ResponseWriter, r *http.Request) {

	nameService := r.FormValue("name")
	typeService := r.FormValue("type")
	newService := api.Service{
		ID:   len(api.GetServices("data.json")) + 1,
		Name: nameService,
		Type: typeService,
	}
	api.AddService("data.json", newService)
}

func AddSlot(w http.ResponseWriter, r *http.Request) {
	serviceidSlot, err := strconv.Atoi(r.FormValue("serviceid"))
	slot := make([]string, 1)
	slot[0] += r.FormValue("slot")
	if err != nil {
		fmt.Println("L'id n'est pas valide !")
		return
	} else {
		api.AddSlotToService("data.json", serviceidSlot, slot)
	}
}

func CancelBooking(w http.ResponseWriter, r *http.Request) {
	bookingId, err := strconv.Atoi(r.FormValue("bookingId"))
	if err != nil {
		fmt.Println("L'id n'est pas valide !")
		return
	} else {
		isBooking := bookingId < len(api.GetBookings("data.json"))
		if isBooking {
			api.RemoveBooking("data.json", GetCookie(r), bookingId)
		} else {
			fmt.Println("Reservation Inexistante !")
			return
		}
	}
}
