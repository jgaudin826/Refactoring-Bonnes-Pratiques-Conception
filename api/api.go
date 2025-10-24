package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

const dataFileName string = "data/data.json"

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

// GetDataJson reads and parses the JSON data file
// Takes the file name as input
// Returns a Data struct containing services, users, and bookings
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
	return result
}
