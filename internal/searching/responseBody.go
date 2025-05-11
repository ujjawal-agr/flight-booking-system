package searching

import "flight-booking-system/internal/enums"

// Flight Response
type FlightResponse struct {
	FlightCode   string                             `json:"flight_code"`
	Company      string                             `json:"company"`
	FlightStatus string                             `json:"flight_status"`
	Seats        map[enums.SeatType]SeatTypeDetails `json:"seats"` // Seat type with price and availability
}

// SeatTypeDetails to hold seat availability and price
type SeatTypeDetails struct {
	Price          int `json:"price"`
	AvailableSeats int `json:"available_seats"`
}
