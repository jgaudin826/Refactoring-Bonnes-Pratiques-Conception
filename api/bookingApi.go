package api

import (
	"encoding/json"
	"fmt"
	"os"
)

// GetBookings retrieves all bookings from the data file
// Returns a slice of Booking structs
func GetBookings() []Booking {
	allData := GetDataJson(dataFileName)
	return allData.Bookings
}

// GetBookingsByEmail retrieves bookings for a specific user by email
// Returns a slice of Booking structs
func GetBookingsByEmail(email string) []Booking {
	dataBookings := GetBookings()
	var userBookings []Booking

	for _, booking := range dataBookings {
		if booking.Email == email {
			userBookings = append(userBookings, booking)
		}
	}
	return userBookings
}

// AddBooking adds a new booking to the data file
func AddBooking(newBooking Booking) {
	allData := GetDataJson(dataFileName)
	allData.Bookings = append(allData.Bookings, newBooking)
	jsonData, errorJsonMarshal := json.Marshal(allData)
	if errorJsonMarshal != nil {
		fmt.Printf("Erreur conversion JSON: %v\n", errorJsonMarshal)
		return
	}

	errorJsonWrite := os.WriteFile(dataFileName, jsonData, 0644)
	if errorJsonWrite != nil {
		fmt.Printf("Erreur écriture fichier: %v\n", errorJsonWrite)
		return
	}
}

// RemoveBooking removes a booking by its ID from the data file
func RemoveBooking(bookingID int) {
	data := GetDataJson(dataFileName)
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

	errWrite := os.WriteFile(dataFileName, jsonData, 0644)
	if errWrite != nil {
		fmt.Println("Erreur écriture fichier:", errWrite)
		return
	}
	AddSlotToService(removedBooking.Service, removedBooking.Slot)
	return
}
