package booking

import (
	"database/sql"
	"errors"
	"flight-booking-system/internal/db"
	"flight-booking-system/internal/enums"
	"fmt"
	"github.com/google/uuid"
)

type Booking struct {
	FlightCode string
	SeatNumber int
	Amount     int
}

func InsertBooking(tx *sql.Tx, flightID uuid.UUID, customerName, customerContact string) (uuid.UUID, error) {
	var bookingID uuid.UUID
	err := tx.QueryRow(`
		INSERT INTO bookings (flight_id, customer_name, customer_contact)
		VALUES ($1, $2, $3) RETURNING booking_id`,
		flightID, customerName, customerContact).Scan(&bookingID)
	if err != nil {
		return uuid.Nil, fmt.Errorf("failed to insert booking: %v", err)
	}
	return bookingID, nil
}

func GetFlightID(tx *sql.Tx, flightCode string) (uuid.UUID, error) {
	var flightID uuid.UUID
	err := tx.QueryRow("SELECT flight_id FROM flights WHERE flight_code = $1", flightCode).Scan(&flightID)
	if err != nil {
		return flightID, fmt.Errorf("failed to fetch flightID: %v", err)
	}
	return flightID, nil
}

func GetAvailableSeat(tx *sql.Tx, flightID uuid.UUID, seatType enums.SeatType) (uuid.UUID, int, error) {
	var seatID uuid.UUID
	var seatNo int

	err := tx.QueryRow(`
		SELECT seat_id, seat_no FROM seats 
		WHERE flight_id = $1 AND seat_type = $2 AND seat_status = 'available'
		ORDER BY seat_no ASC 
		LIMIT 1`,
		flightID, seatType).Scan(&seatID, &seatNo)

	if errors.Is(err, sql.ErrNoRows) {
		return uuid.Nil, 0, fmt.Errorf("no available seat of type %s", seatType)
	} else if err != nil {
		return uuid.Nil, 0, fmt.Errorf("failed to find seat: %v", err)
	}

	return seatID, seatNo, nil
}

func BookSeat(tx *sql.Tx, seatID uuid.UUID) error {
	_, err := tx.Exec(`UPDATE seats SET seat_status = 'booked' WHERE seat_id = $1`, seatID)
	return err
}

func InsertBookingSeatMapping(tx *sql.Tx, bookingID, seatID uuid.UUID, name string, age int, gender enums.Gender) error {
	_, err := tx.Exec(`
		INSERT INTO booking_seat_mapping (booking_id, seat_id, name, age, gender)
		VALUES ($1, $2, $3, $4, $5)`,
		bookingID, seatID, name, age, gender)
	return err
}

func GetSeatPrice(tx *sql.Tx, flightID uuid.UUID, seatType enums.SeatType) (int, error) {
	var price int
	err := tx.QueryRow(`
		SELECT price FROM pricing
		WHERE flight_id = $1 AND seat_type = $2`,
		flightID, seatType).Scan(&price)
	if err != nil {
		return 0, fmt.Errorf("failed to fetch price: %v", err)
	}
	return price, nil
}

func CancelBooking(bookingID uuid.UUID) error {
	conn := db.GetDB()

	// fetch flight_id for incrementing seats
	var flightID uuid.UUID
	err := conn.QueryRow("SELECT flight_id FROM bookings WHERE booking_id = $1", bookingID).Scan(&flightID)
	if err != nil {
		return err
	}

	_, err = conn.Exec("DELETE FROM bookings WHERE booking_id = $1", bookingID)
	if err != nil {
		return err
	}

	// You might want to update seat status back to 'available' here if needed.
	return nil
}
