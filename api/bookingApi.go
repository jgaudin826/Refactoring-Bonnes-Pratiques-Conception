package api

import (
	"encoding/json"
	"fmt"
	"os"
)

func GetBookings(fileName string) []Booking {
	allData := GetDataJson(fileName)
	return allData.Bookings
}

func AddBooking(fileName string, newBooking Booking) {
	allData := GetDataJson(fileName)
	allData.Bookings = append(allData.Bookings, newBooking)
	jsonData, errorJsonMarshal := json.Marshal(allData)
	if errorJsonMarshal != nil {
		fmt.Printf("Erreur conversion JSON: %v\n", errorJsonMarshal, "\n")
		return
	}

	errorJsonWrite := os.WriteFile(fileName, jsonData, 0644)
	if errorJsonWrite != nil {
		fmt.Printf("Erreur écriture fichier: %v\n", errorJsonWrite, "\n")
		return
	}
}

func RemoveBooking(fileName string, bookingID int) {
	data := GetDataJson(fileName)
	var updatedBookings []Booking
	var removedBooking Booking
	found := false

	for _, bookings := range data.Bookings {
		if bookings.ID == bookingID {
			found = true
			removedBooking = bookings
			continue
		}
		updatedBookings = append(updatedBookings, bookings)

	}

	if !found {
		fmt.Printf("Réservation #%d introuvable ou non autorisée pour %s.\n", bookingID, "\n")
		return
	}
	data.Bookings = updatedBookings
	jsonData, errMarshal := json.MarshalIndent(data, "", "  ")
	if errMarshal != nil {
		fmt.Println("Erreur conversion JSON:", errMarshal)
		return
	}

	errWrite := os.WriteFile(fileName, jsonData, 0644)
	if errWrite != nil {
		fmt.Println("Erreur écriture fichier:", errWrite)
		return
	}
	AddSlotToService(fileName, removedBooking.Service, removedBooking.Slot)
	return
}
