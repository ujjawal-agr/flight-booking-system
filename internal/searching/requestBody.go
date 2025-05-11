package searching

import "time"

type GetFlightsRequest struct {
	Source      string    `json:"source"`
	Destination string    `json:"destination"`
	Date        time.Time `json:"date"`
}
