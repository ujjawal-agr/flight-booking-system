package scheduling

import "net/http"

func RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/scheduleFlight", ScheduleFlightHandler)
}
