package services

import (
	"fmt"
	"net/http"
	"refactoring/api"
	"strconv"
	"time"
)

// Create Service from user submitted form and add it to the dataFile by calling AddService
func CreateService(write http.ResponseWriter, request *http.Request) {

	nameService := request.FormValue("name")
	typeService := request.FormValue("type")
	newService := api.Service{
		ID:   len(api.GetServices()) + 1,
		Name: nameService,
		Type: typeService,
	}
	api.AddService(newService)
	http.Redirect(write, request, "/", http.StatusSeeOther)
}

// Create Slot from user submitted form and add it to the dataFile by calling AddSlotToService
func CreateSlot(write http.ResponseWriter, request *http.Request) {
	serviceidSlot, err := strconv.Atoi(request.FormValue("serviceid"))
	slot, errorParse := time.Parse("2006-01-02T15:04", request.FormValue("slot"))
	if err != nil || errorParse != nil {
		fmt.Println("L'id ou le slot n'est pas valide !")
		http.Redirect(write, request, "/", http.StatusSeeOther)
		return
	} else {
		api.AddSlotToService(serviceidSlot, slot.Format("2006-01-02 15:04"))
		http.Redirect(write, request, "/", http.StatusSeeOther)
	}
}

// Cancel booking by clicking a form button and calling RemoveBooking
func CancelBooking(write http.ResponseWriter, request *http.Request) {
	bookingId, err := strconv.Atoi(request.FormValue("bookingId"))
	if err != nil {
		fmt.Println("L'id n'est pas valide !")
		http.Redirect(write, request, "/", http.StatusSeeOther)
	} else {
		api.RemoveBooking(bookingId)
		http.Redirect(write, request, "/", http.StatusSeeOther)
	}
}
