package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
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
		fmt.Println("Erreur lecture fichier:", errorJsonRead, "\n")
		return Data{}
	}

	var result Data
	errorParsing := json.Unmarshal(allData, &result)
	if errorParsing != nil {
		fmt.Println("Erreur parsing JSON:", errorParsing, "\n")
		return Data{}
	}
	return result
}
