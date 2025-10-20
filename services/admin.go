package services

import (
	"net/http"
	"refactoring/api"
)

func AddService(w http.ResponseWriter, r *http.Request) {
	nameService := r.FormValue("name")
	typeService := r.FormValue("type")

}

func AddSlot(w http.ResponseWriter, r *http.Request) {
	serviceidSlot := r.FormValue("serviceid")
	slot := r.FormValue("slot")
	api.AddSlotToService("data.json", serviceidSlot, slot)
}

func CancelBooking(w http.ResponseWriter, r *http.Request) {
	bookingId := r.FormValue("bookingId")

}
