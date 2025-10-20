package services

import (
	"fmt"
	"refactoring/api"
)

func BookingSlot(fileName, email string, serviceID int, slot string) {
	dataBookings := api.GetBookings(fileName)
	dataServices := api.GetServices(fileName)

	var service *api.Service
	for index, services := range dataServices {
		if services.ID == serviceID {
			service = &dataServices[index]
			break
		}
	}
	if service == nil {
		fmt.Println("Service introuvable.")
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
		return
	}

	for _, bookings := range dataBookings {
		if bookings.Email == email && bookings.Service == serviceID && bookings.Slot == slot {
			fmt.Println("Vous avez déjà réservé ce créneau.")
			return
		}
	}

	for _, bookings := range dataBookings {
		if bookings.Service == serviceID && bookings.Slot == slot {
			fmt.Println("Ce créneau est déjà complet.")
			return
		}
	}
	newBooking := api.Booking{
		ID:      len(dataBookings) + 1,
		Email:   email,
		Service: serviceID,
		Slot:    slot,
	}
	api.AddBooking(fileName, newBooking)
	api.RemoveSlotFromService(fileName, newBooking.ID, []string{slot})
	fmt.Println("Réservation réussie.")
	return
}

func GetBookingsByEmail(fileName, email string) []api.Booking {
	dataBookings := api.GetBookings(fileName)
	var userBookings []api.Booking

	for _, booking := range dataBookings {
		if booking.Email == email {
			userBookings = append(userBookings, booking)
		}
	}
	return userBookings
}
