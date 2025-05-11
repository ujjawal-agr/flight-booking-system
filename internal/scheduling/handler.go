package scheduling

import (
	"encoding/json"
	"net/http"
)

func ScheduleFlightHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST allowed", http.StatusMethodNotAllowed)
		return
	}

	var req ScheduleFlightRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	err = ScheduleFlightController(req)
	if err != nil {
		http.Error(w, "Failed to schedule flight: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Flight scheduled successfully"))
}
