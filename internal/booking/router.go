package booking

import "net/http"

func RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/bookFlight", BookFlightHandler)
	mux.HandleFunc("/cancelBooking", CancelBookingHandler)
}
