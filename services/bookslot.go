package services

import (
	"fmt"
	"refactoring/api"
)

func BookSlot(fileName, email string, serviceID int, slot string) {
	dataBookings := api.GetBookings(fileName)

	dataServices := api.GetServices(fileName)

	var service *api.Service
	for i, s := range dataServices {
		if s.ID == serviceID {
			service = &dataServices[i]
			break
		}
	}
	if service == nil {
		fmt.Println("❌ Service introuvable.")
		return
	}

	slotExists := false
	for _, s := range service.Slots {
		if s == slot {
			slotExists = true
			break
		}
	}
	if !slotExists {
		fmt.Println("❌ Ce créneau n'existe pas pour ce service.")
		return
	}

	for _, b := range dataBookings {
		if b.Email == email && b.Service == serviceID && b.Slot == slot {
			fmt.Println("⚠️ Vous avez déjà réservé ce créneau.")
			return
		}
	}

	for _, b := range dataBookings {
		if b.Service == serviceID && b.Slot == slot {
			fmt.Println("⚠️ Ce créneau est déjà complet.")
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
}
