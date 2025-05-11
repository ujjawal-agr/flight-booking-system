package booking

import (
	"flight-booking-system/internal/enums"
)

type Seat struct {
	SeatType enums.SeatType `json:"seat_type"`
	Name     string         `json:"name"`
	Age      int            `json:"age"`
	Gender   enums.Gender   `json:"gender"`
}
type BookingRequest struct {
	FlightCode      string `json:"flight_code"`
	CustomerName    string `json:"customer_name"`
	CustomerContact string `json:"customer_contact"`
	SeatInfo        []Seat `json:"seat_info"`
}
