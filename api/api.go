package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Service struct {
	ID    int      `json:"id"`
	Name  string   `json:"name"`
	Type  string   `json:"type"`
	Slots []string `json:"slots"`
}

type User struct {
	ID    int    `json:"id"`
	Email string `json:"email"`
	Role  string `json:"role"`
}

type Booking struct {
	ID      int    `json:"id"`
	Email   string `json:"email"`
	Service int    `json:"service"`
	Slot    string `json:"slot"`
}

type Data struct {
	Services []Service `json:"services"`
	Users    []User    `json:"users"`
	Bookings []Booking `json:"bookings"`
}

func GetDataJson(fileName string) Data {
	allData, errorJsonRead := ioutil.ReadFile(fileName)
	if errorJsonRead != nil {
		fmt.Println("Erreur lecture fichier:", errorJsonRead)
		return Data{}
	}

	var result Data
	errorParsing := json.Unmarshal(allData, &result)
	if errorParsing != nil {
		fmt.Println("Erreur parsing JSON:", errorParsing)
		return Data{}
	}

	fmt.Printf("%+v\n", result)
	return result
}

func GetUsers(fileName string) []User {
	allData := GetDataJson(fileName)
	return allData.Users
}

func GetServices(fileName string) []Service {
	allData := GetDataJson(fileName)
	return allData.Services
}

func GetBookings(fileName string) []Booking {
	allData := GetDataJson(fileName)
	return allData.Bookings
}

func AddBooking(fileName string, newBooking Booking) {
	allData := GetDataJson(fileName)
	allData.Bookings = append(allData.Bookings, newBooking)

	jsonData, errorJsonMarshal := json.Marshal(allData)
	if errorJsonMarshal != nil {
		fmt.Printf("Erreur conversion JSON: %v\n", errorJsonMarshal)
		return
	}

	errorJsonWrite := os.WriteFile(fileName, jsonData, 0644)
	if errorJsonWrite != nil {
		fmt.Printf("Erreur écriture fichier: %v\n", errorJsonWrite)
		return
	}
}

func AddService(fileName string, newService Service) {
	allData := GetDataJson(fileName)
	allData.Services = append(allData.Services, newService)

	jsonData, errorJsonMarshal := json.Marshal(allData)
	if errorJsonMarshal != nil {
		fmt.Printf("Erreur conversion JSON: %v\n", errorJsonMarshal)
		return
	}

	errorJsonWrite := os.WriteFile(fileName, jsonData, 0644)
	if errorJsonWrite != nil {
		fmt.Printf("Erreur écriture fichier: %v\n", errorJsonWrite)
		return
	}
}

func AddUser(fileName string, newUser User) {
	allData := GetDataJson(fileName)
	allData.Users = append(allData.Users, newUser)

	jsonData, errorJsonMarshal := json.Marshal(allData)
	if errorJsonMarshal != nil {
		fmt.Printf("Erreur conversion JSON: %v\n", errorJsonMarshal)
		return
	}

	errorJsonWrite := os.WriteFile(fileName, jsonData, 0644)
	if errorJsonWrite != nil {
		fmt.Printf("Erreur écriture fichier: %v\n", errorJsonWrite)
		return
	}
}

func AddSlotToService(fileName string, serviceID int, newSlots []string) {
	allData := GetDataJson(fileName)
	found := false

	for index, service := range allData.Services {
		if service.ID == serviceID {
			allData.Services[index].Slots = append(allData.Services[index].Slots, newSlots...)
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

	errorJsonWrite := os.WriteFile(fileName, jsonData, 0644)
	if errorJsonWrite != nil {
		fmt.Printf("Erreur écriture fichier: %v\n", errorJsonWrite)
		return
	}

	fmt.Printf("Slots ajoutés au service %d avec succès !\n", serviceID)
}

func RemoveSlotFromService(fileName string, serviceID int, slotsToRemove []string) {
	allData := GetDataJson(fileName)
	found := false

	for i, service := range allData.Services {
		if service.ID == serviceID {
			found = true
			var updatedSlots []string
			for _, slot := range service.Slots {
				shouldRemove := false
				for _, slot := range slotsToRemove {
					if slot == slot {
						shouldRemove = true
						break
					}
				}
				if !shouldRemove {
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

	jsonData, errorJsonMarshal := json.MarshalIndent(allData, "", "  ")
	if errorJsonMarshal != nil {
		fmt.Printf("Erreur conversion JSON: %v\n", errorJsonMarshal)
		return
	}

	errorJsonWrite := os.WriteFile(fileName, jsonData, 0644)
	if errorJsonWrite != nil {
		fmt.Printf("Erreur écriture fichier: %v\n", errorJsonWrite)
		return
	}
}

func RemoveBooking(fileName, email string, bookingID int) {
	data := GetDataJson(fileName)
	var updatedBookings []Booking
	found := false

	for _, bookings := range data.Bookings {
		if bookings.ID == bookingID && bookings.Email == email {
			found = true
			continue
		}
		updatedBookings = append(updatedBookings, bookings)
	}

	if !found {
		fmt.Printf("Réservation #%d introuvable ou non autorisée pour %s.\n", bookingID, email)
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

}
