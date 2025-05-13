package scheduling

import (
	"flight-booking-system/internal/enums"
	"time"
)

type ScheduleFlightRequest struct {
	FlightCode  string                 `json:"flight_code"`
	Source      string                 `json:"source"`
	Destination string                 `json:"destination"`
	Date        time.Time              `json:"date"`
	Company     string                 `json:"company"`
	SeatCounts  map[enums.SeatType]int `json:"seat_counts"`
	Prices      map[enums.SeatType]int `json:"prices"`
	CreatedBy   string                 `json:"created_by"`
}

type CancelFlightRequest struct {
	FlightCode string `json:"flight_code"`
}
