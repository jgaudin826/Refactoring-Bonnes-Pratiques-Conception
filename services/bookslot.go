package services

import (
	"encoding/json"
	"fmt"
	"os"
	"refactoring/api"
)

func BookSlot(fileName, email string, serviceID int, slot string) {
	data := api.GetDataJson(fileName)

	services := api.GetServices(fileName)

	var service *api.Service
	for i, s := range services {
		if s.ID == serviceID {
			service = &services[i]
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

	for _, b := range data.Bookings {
		if b.Email == email && b.Service == serviceID && b.Slot == slot {
			fmt.Println("⚠️ Vous avez déjà réservé ce créneau.")
			return
		}
	}

	for _, b := range data.Bookings {
		if b.Service == serviceID && b.Slot == slot {
			fmt.Println("⚠️ Ce créneau est déjà complet.")
			return
		}
	}

	newBooking := api.Booking{
		ID:      len(data.Bookings) + 1,
		Email:   email,
		Service: serviceID,
		Slot:    slot,
	}
	data.Bookings = append(data.Bookings, newBooking)

	jsonData, errorJsonMarshal := json.MarshalIndent(data, "", "  ")
	if errorJsonMarshal != nil {
		fmt.Println("Erreur conversion JSON:", errorJsonMarshal)
		return
	}

	errorJsonWrite := os.WriteFile(fileName, jsonData, 0644)
	if errorJsonWrite != nil {
		fmt.Println("Erreur écriture fichier:", errorJsonWrite)
		return
	}

	fmt.Printf("✅ Réservation confirmée pour %s à %s sur le service %s\n", email, slot, service.Name)
}
