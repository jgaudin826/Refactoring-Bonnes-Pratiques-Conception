package api

import (
	"encoding/json"
	"fmt"
	"os"
)

func GetUsers(fileName string) []User {
	allData := GetDataJson(fileName)
	return allData.Users
}

func AddUser(fileName string, newUser User) {
	allData := GetDataJson(fileName)
	allData.Users = append(allData.Users, newUser)

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
