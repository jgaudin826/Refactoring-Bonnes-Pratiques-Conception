package services

import (
	"fmt"
	"net/http"
	"refactoring/api"
	"strconv"
	"time"
)

func AddService(w http.ResponseWriter, r *http.Request) {

	nameService := r.FormValue("name")
	typeService := r.FormValue("type")
	newService := api.Service{
		ID:   len(api.GetServices("data/data.json")) + 1,
		Name: nameService,
		Type: typeService,
	}
	api.AddService("data/data.json", newService)
	http.Redirect(w, r, "/", http.StatusSeeOther)
	return
}

func AddSlot(w http.ResponseWriter, r *http.Request) {
	serviceidSlot, err := strconv.Atoi(r.FormValue("serviceid"))
	slot, errorParse := time.Parse("2006-01-02T15:04", r.FormValue("slot"))
	if err != nil || errorParse != nil {
		fmt.Println("L'id ou le slot n'est pas valide !")
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	} else {
		api.AddSlotToService("data/data.json", serviceidSlot, slot.Format("2006-01-02 15:04"))
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
}

func CancelBooking(w http.ResponseWriter, r *http.Request) {
	bookingId, err := strconv.Atoi(r.FormValue("bookingId"))
	if err != nil {
		fmt.Println("L'id n'est pas valide !")
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	} else {
		api.RemoveBooking("data/data.json", bookingId)
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
}
