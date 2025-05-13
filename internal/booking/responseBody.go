package booking

import "flight-booking-system/internal/enums"

type Bill struct {
	PassengerName string         `json:"passenger_name"`
	FlightCode    string         `json:"flight_code"`
	SeatType      enums.SeatType `json:"seat_type"`
	SeatNo        int            `json:"seat_no"`
	Amount        int            `json:"amount"`
}
