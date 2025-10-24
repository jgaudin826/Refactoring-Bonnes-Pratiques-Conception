package api

import (
	"encoding/json"
	"fmt"
	"os"
)

// GetServices retrieves all services from the data file
// Returns a slice of Service structs
func GetServices() []Service {
	allData := GetDataJson(dataFileName)
	return allData.Services
}

// AddService adds a new service to the data file
func AddService(newService Service) {
	allData := GetDataJson(dataFileName)
	allData.Services = append(allData.Services, newService)

	jsonData, errorJsonMarshalIndent := json.MarshalIndent(allData, "", "  ")
	if errorJsonMarshalIndent != nil {
		fmt.Printf("Erreur conversion JSON: %v\n", errorJsonMarshalIndent)
		return
	}

	errorJsonWrite := os.WriteFile(dataFileName, jsonData, 0644)
	if errorJsonWrite != nil {
		fmt.Printf("Erreur écriture fichier: %v\n", errorJsonWrite)
		return
	}
}
