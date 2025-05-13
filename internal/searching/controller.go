package searching

import (
	"flight-booking-system/internal/enums"
	"fmt"
)

// GetFlights to get flights based on source, destination, and date
func GetFlights(req GetFlightsRequest) ([]FlightResponse, error) {

	// Fetch flights from the database using the repository
	flights, err := FetchFlights(req)
	if err != nil {
		return nil, err
	}

	seatTypes := []string{"window", "middle", "aisle"}
	// Process each flight and determine flight status based on seat availability
	for i, flight := range flights {
		// Get available seats and prices for each seat type
		var seatsAvailable = make(map[string]int)
		var seatsPrice = make(map[string]int)
		flight.Seats = make(map[enums.SeatType]SeatTypeDetails)
		if flight.FlightStatus == "cancelled" {
			for _, seatType := range seatTypes {
				seatsAvailable[seatType] = 0
				seatsPrice[seatType] = 0
				flight.Seats[enums.SeatType(seatType)] = SeatTypeDetails{Price: seatsPrice[seatType], AvailableSeats: seatsAvailable[seatType]}
			}
		} else {
			flightID, err := GetFlightID(flight.FlightCode)
			if err != nil {
				return nil, err
			}
			for _, seatType := range seatTypes {
				seatsAvailable[seatType], err = CountAvailableSeats(flightID, seatType)
				if err != nil {
					return nil, fmt.Errorf("failed to count available seats: %v", err)
				}
				seatsPrice[seatType], err = GetSeatPrice(flightID, seatType)
				if err != nil {
					return nil, fmt.Errorf("failed to fetch seat prices: %v", err)
				}

				flight.Seats[enums.SeatType(seatType)] = SeatTypeDetails{Price: seatsPrice[seatType], AvailableSeats: seatsAvailable[seatType]}
			}

			// Determine flight status based on available seats
			if seatsAvailable[seatTypes[0]] == 0 && seatsAvailable[seatTypes[1]] == 0 && seatsAvailable[seatTypes[2]] == 0 {
				flight.FlightStatus = "fullyBooked"
				err := SetFullyBooked(flightID)
				if err != nil {
					return nil, err
				}
				//else {
				//	_ = cache.SetFlightAsFullyBooked(flight.FlightCode, flight.Date.Format("2006-01-02"))
				//}
			} else {
				flight.FlightStatus = "available"
			}
		}
		// Update the flight in the list
		flights[i] = flight
	}

	return flights, nil
}
