package utils

import (
	"encoding/json"
	"net/http"
)

// ValidateInput checks if the input is valid.
func ValidateInput(input interface{}) error {
	// Implement validation logic here
	return nil
}

// FormatResponse formats the response to be sent to the client.
func FormatResponse(w http.ResponseWriter, data interface{}, status int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}