package searching

import (
	"flight-booking-system/internal/db"
	"fmt"
	"github.com/google/uuid"
)

// FetchFlights to fetch flight data based on source, destination, and date
func FetchFlights(req GetFlightsRequest) ([]FlightResponse, error) {
	conn := db.GetDB()

	// Query flights based on source, destination, and date
	rows, err := conn.Query(`
		SELECT flight_code, company, flight_status
		FROM flights
		WHERE source = $1 AND destination = $2 AND date = $3`,
		req.Source, req.Destination, req.Date)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch flights: %v", err)
	}
	defer rows.Close()

	var flights []FlightResponse

	// Iterate through each flight
	for rows.Next() {
		var flightCode, company, flightStatus string
		err := rows.Scan(&flightCode, &company, &flightStatus)
		if err != nil {
			return nil, fmt.Errorf("failed to scan flight row: %v", err)
		}

		// Create the flight response object
		flight := FlightResponse{
			FlightCode:   flightCode,
			Company:      company,
			FlightStatus: flightStatus,
		}
		flights = append(flights, flight)
	}

	return flights, nil
}

func GetFlightID(flightCode string) (uuid.UUID, error) {
	conn := db.GetDB()
	var flightID uuid.UUID
	err := conn.QueryRow("SELECT flight_id FROM flights WHERE flight_code = $1", flightCode).Scan(&flightID)
	if err != nil {
		return flightID, fmt.Errorf("failed to fetch flightID: %v", err)
	}
	return flightID, nil
}

// CountAvailableSeats to get available seats by seat type
func CountAvailableSeats(flightID uuid.UUID, seatType string) (int, error) {
	conn := db.GetDB()

	// Query to count available seats for each seat type
	//var seatsAvailable SeatsAvailable
	var seatsNumber int
	// Count window seats
	err := conn.QueryRow(`
		SELECT COUNT(*) FROM seats WHERE flight_id = $1 AND seat_type = $2 AND seat_status = 'available'`,
		flightID, seatType).Scan(&seatsNumber)
	if err != nil {
		return seatsNumber, fmt.Errorf("failed to count seats: %v", err)
	}
	return seatsNumber, nil
}

// GetSeatPrices to fetch the price for each seat type
func GetSeatPrice(flightID uuid.UUID, seatType string) (int, error) {
	conn := db.GetDB()

	// Query to get the price for each seat type
	var price int

	// Get price for window seat
	err := conn.QueryRow(`
		SELECT price FROM pricing WHERE flight_id = $1 AND seat_type = $2`,
		flightID, seatType).Scan(&price)
	if err != nil {
		return price, fmt.Errorf("failed to fetch window seat price: %v", err)
	}
	return price, nil
}

func SetFullyBooked(flightID uuid.UUID) error {
	conn := db.GetDB()
	_, err := conn.Exec(`UPDATE flights SET flight_status = 'fullyBooked' WHERE flight_id = $1`, flightID)
	return err
}
