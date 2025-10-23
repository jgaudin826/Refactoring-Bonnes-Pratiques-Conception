package services

import (
	"fmt"
	"net/http"
	"refactoring/api"
	"strconv"
	"time"
)

func AddService(write http.ResponseWriter, request *http.Request) {

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

func AddSlot(write http.ResponseWriter, request *http.Request) {
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
