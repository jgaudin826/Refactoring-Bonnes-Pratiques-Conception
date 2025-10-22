package api

import (
	"encoding/json"
	"fmt"
	"os"
)

func GetServices(fileName string) []Service {
	allData := GetDataJson(fileName)
	return allData.Services
}

func AddService(fileName string, newService Service) {
	allData := GetDataJson(fileName)
	allData.Services = append(allData.Services, newService)

	jsonData, errorJsonMarshalIndent := json.MarshalIndent(allData, "", "  ")
	if errorJsonMarshalIndent != nil {
		fmt.Printf("Erreur conversion JSON: %v\n", errorJsonMarshalIndent, "\n")
		return
	}

	errorJsonWrite := os.WriteFile(fileName, jsonData, 0644)
	if errorJsonWrite != nil {
		fmt.Printf("Erreur écriture fichier: %v\n", errorJsonWrite, "\n")
		return
	}
}
