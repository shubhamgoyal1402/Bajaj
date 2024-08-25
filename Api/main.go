package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
)

// Request structure to hold input JSON data
type Request struct {
	Data []string `json:"data"`
}

// Response structure to hold the filtered response data
type Response struct {
	IsSuccess                bool     `json:"is_success"`
	UserID                   string   `json:"user_id"`
	Email                    string   `json:"email"`
	RollNumber               string   `json:"roll_number"`
	Numbers                  []string `json:"numbers"`
	Alphabets                []string `json:"alphabets"`
	HighestLowercaseAlphabet string   `json:"highest_lowercase_alphabet"`
}

// HardcodedResponse structure for GET method
type HardcodedResponse struct {
	OperationCode int `json:"operation_code"`
}

// Handler function to process incoming requests
func processRequest(w http.ResponseWriter, r *http.Request) {
	// Set the content type of the response
	w.Header().Set("Content-Type", "application/json")

	// Check the HTTP method
	if r.Method == http.MethodGet {
		// Handle GET request
		hardcodedResponse := HardcodedResponse{
			OperationCode: 1,
		}
		jsonResponse, err := json.Marshal(hardcodedResponse)
		if err != nil {
			http.Error(w, "Failed to encode response as JSON", http.StatusInternalServerError)
			return
		}
		w.Write(jsonResponse)
	} else if r.Method == http.MethodPost {
		// Handle POST request
		var requestData Request
		err := json.NewDecoder(r.Body).Decode(&requestData)
		if err != nil {
			http.Error(w, "Invalid JSON format", http.StatusBadRequest)
			return
		}

		// Initialize response object with user details
		response := Response{
			IsSuccess:  true,
			UserID:     "tisha_mahajan_29112002",   // Replace with actual user id
			Email:      "tishamahajan43@gmail.com", // Replace with actual email
			RollNumber: "21BBS0086",                // Replace with actual roll number
		}

		// Initialize data structures to hold numbers and alphabets
		var numbers []string
		var alphabets []string
		highestLowercase := ""

		// Iterate through input data to separate numbers and alphabets
		for _, item := range requestData.Data {
			if _, err := strconv.Atoi(item); err == nil {
				// Item is a number
				numbers = append(numbers, item)
			} else if isAlphabet(item) {
				// Item is an alphabet
				alphabets = append(alphabets, item)
				// Check for highest lowercase alphabet
				if item == strings.ToLower(item) && (highestLowercase == "" || item > highestLowercase) {
					highestLowercase = item
				}
			}
		}

		// Populate the response object with filtered data
		response.Numbers = numbers
		response.Alphabets = alphabets
		response.HighestLowercaseAlphabet = highestLowercase

		// Send the response as JSON
		jsonResponse, err := json.Marshal(response)
		if err != nil {
			http.Error(w, "Failed to encode response as JSON", http.StatusInternalServerError)
			return
		}
		w.Write(jsonResponse)
	} else {
		// Unsupported method
		http.Error(w, "Unsupported method", http.StatusMethodNotAllowed)
	}
}

// Utility function to check if a string is an alphabet
func isAlphabet(s string) bool {
	s = strings.ToLower(s)
	return s >= "a" && s <= "z"
}

// Main function to start the HTTP server
func main() {
	fileServer := http.FileServer(http.Dir("./static"))

	http.Handle("/", fileServer)

	http.HandleFunc("/bfhl", processRequest) // Single endpoint handling both GET and POST requests

	// Start the server on port 8080
	fmt.Println("Server is running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
