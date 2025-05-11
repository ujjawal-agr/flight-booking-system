package searching

import "net/http"

func RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/searchFlight", GetFlightsHandler)
}
