package api

import (
	"encoding/json"
	"fmt"
	"os"
)

func GetUsers() []User {
	allData := GetDataJson(dataFileName)
	return allData.Users
}

func AddUser(newUser User) {
	allData := GetDataJson(dataFileName)
	allData.Users = append(allData.Users, newUser)

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
