package api

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

// AddSlotToService adds a new slot to a service by its ID
// Check if the service already exists
// Updates the data file with the new slot
func AddSlotToService(serviceID int, newSlots string) {
	allData := GetDataJson(dataFileName)
	found := false

	for index, service := range allData.Services {
		if service.ID == serviceID {
			allData.Services[index].Slots = append(allData.Services[index].Slots, newSlots)
			found = true
			break
		}
	}

	if !found {
		fmt.Printf("Service avec ID %d non trouvé.\n", serviceID)
		return
	}

	jsonData, errorJsonMarshal := json.MarshalIndent(allData, "", "  ")
	if errorJsonMarshal != nil {
		fmt.Printf("Erreur conversion JSON: %v\n", errorJsonMarshal)
		return
	}

	errorJsonWrite := os.WriteFile(dataFileName, jsonData, 0644)
	if errorJsonWrite != nil {
		fmt.Printf("Erreur écriture fichier: %v\n", errorJsonWrite)
		return
	}
	return
}

// RemoveSlotFromService removes a slot from a service by its ID
// checks if the service exists and updates the data file
func RemoveSlotFromService(serviceID int, slotToRemove string) {
	allData := GetDataJson(dataFileName)
	found := false

	for i, service := range allData.Services {
		if service.ID == serviceID {
			found = true
			var updatedSlots []string
			for _, slot := range service.Slots {
				if strings.TrimSpace(slot) != strings.TrimSpace(slotToRemove) {
					updatedSlots = append(updatedSlots, slot)
				}
			}

			allData.Services[i].Slots = updatedSlots
			break
		}
	}

	if !found {
		fmt.Printf("Service avec ID %d non trouvé.\n", serviceID)
		return
	}

	jsonData, err := json.MarshalIndent(allData, "", "  ")
	if err != nil {
		fmt.Printf("Erreur conversion JSON: %v\n", err)
		return
	}

	if err := os.WriteFile(dataFileName, jsonData, 0644); err != nil {
		fmt.Printf("Erreur écriture fichier: %v\n", err)
		return
	}
}
