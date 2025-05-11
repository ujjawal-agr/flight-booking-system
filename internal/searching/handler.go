package searching

import (
	"encoding/json"
	"net/http"
)

func GetFlightsHandler(w http.ResponseWriter, r *http.Request) {
	var req GetFlightsRequest
	// Parse request body
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	// Fetch available flights from controller
	flights, err := GetFlights(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Return the flight details as JSON response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(flights)
}
