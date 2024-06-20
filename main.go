package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type EntityResponse struct {
	IsValid    bool   `json:"isValid"`
	UEN        string `json:"uen"`
	Name       string `json:"name"`
	EntityType string `json:"entityType"`
}

func main() {
	http.HandleFunc("/entityVerification", func(w http.ResponseWriter, r *http.Request) {
		// Check for the "token" header
		token := r.Header.Get("token")
		if token != "5577" {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		// Parse query parameters
		uen := r.URL.Query().Get("uen")

		// Validate the UEN or fetch corresponding data (mocking for example)
		// Here, we use a simple mock response based on the UEN parameter
		response := getEntityData(uen)

		// Marshal response to JSON
		jsonResponse, err := json.Marshal(response)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Set content type and write response
		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonResponse)
	})

	// Start the server on port 8080
	log.Println("Starting server on port 8080...")
	err := http.ListenAndServe(":3300", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

// Mock function to return data based on UEN (in real application, this would fetch from database or service)
func getEntityData(uen string) EntityResponse {
	// For demonstration purposes, return different mock data based on UEN
	switch uen {
	case "999999999N":
		return EntityResponse{
			IsValid:    true,
			UEN:        "999999999N",
			Name:       "ABC Company",
			EntityType: "COM",
		}
	case "123456789A":
		return EntityResponse{
			IsValid:    true,
			UEN:        "123456789A",
			Name:       "XYZ Enterprises",
			EntityType: "COM",
		}
	case "987654321B":
		return EntityResponse{
			IsValid:    true,
			UEN:        "987654321B",
			Name:       "DEF Industries",
			EntityType: "COM",
		}
	case "111222333C":
		return EntityResponse{
			IsValid:    true,
			UEN:        "111222333C",
			Name:       "GHI Holdings",
			EntityType: "COM",
		}
	case "444555666D":
		return EntityResponse{
			IsValid:    true,
			UEN:        "444555666D",
			Name:       "JKL Corporation",
			EntityType: "COM",
		}
	case "777888999E":
		return EntityResponse{
			IsValid:    true,
			UEN:        "777888999E",
			Name:       "MNO Ventures",
			EntityType: "COM",
		}
	case "222333444F":
		return EntityResponse{
			IsValid:    true,
			UEN:        "222333444F",
			Name:       "PQR Ltd",
			EntityType: "COM",
		}
	case "555666777G":
		return EntityResponse{
			IsValid:    true,
			UEN:        "555666777G",
			Name:       "STU Group",
			EntityType: "COM",
		}
	case "888999111H":
		return EntityResponse{
			IsValid:    true,
			UEN:        "888999111H",
			Name:       "VWX Partners",
			EntityType: "COM",
		}
	case "333444555I":
		return EntityResponse{
			IsValid:    true,
			UEN:        "333444555I",
			Name:       "YZA Enterprises",
			EntityType: "COM",
		}
	case "666777888J":
		return EntityResponse{
			IsValid:    true,
			UEN:        "666777888J",
			Name:       "BCD Solutions",
			EntityType: "COM",
		}
	case "999111222K":
		return EntityResponse{
			IsValid:    true,
			UEN:        "999111222K",
			Name:       "EFG Innovations",
			EntityType: "COM",
		}
	case "444555666L":
		return EntityResponse{
			IsValid:    true,
			UEN:        "444555666L",
			Name:       "HIJ Systems",
			EntityType: "COM",
		}
	case "777888999M":
		return EntityResponse{
			IsValid:    true,
			UEN:        "777888999M",
			Name:       "KLM Technologies",
			EntityType: "COM",
		}
	case "222333444N":
		return EntityResponse{
			IsValid:    true,
			UEN:        "222333444N",
			Name:       "NOP Enterprises",
			EntityType: "COM",
		}
	case "555666777O":
		return EntityResponse{
			IsValid:    true,
			UEN:        "555666777O",
			Name:       "QRS Holdings",
			EntityType: "COM",
		}
	default:
		return EntityResponse{
			IsValid: false,
		}
	}
}
