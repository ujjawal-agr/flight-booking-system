package scheduling

import (
	"database/sql"
	"flight-booking-system/internal/enums"
	"github.com/google/uuid"
	"time"
)

func InsertFlight(tx *sql.Tx, flightID uuid.UUID, req ScheduleFlightRequest) error {

	_, err := tx.Exec(`
		INSERT INTO flights (flight_id, flight_code, source, destination, company, flight_status, date, created_on, created_by)
		VALUES ($1, $2, $3, $4, $5, 'available', $6, $7, $8)
	`,
		flightID, req.FlightCode, req.Source, req.Destination,
		req.Company, req.Date, time.Now(), req.CreatedBy,
	)
	return err
}

func InsertSeat(tx *sql.Tx, seatID uuid.UUID, flightID uuid.UUID, seatNo int, seatType enums.SeatType) error {
	//fmt.Printf("hello")

	_, err := tx.Exec(`
		INSERT INTO seats (seat_id, flight_id, seat_no, seat_type, seat_status, created_on)
		VALUES ($1, $2, $3, $4, 'available', $5)
	`,
		seatID, flightID, seatNo, seatType, time.Now(),
	)
	return err
}

func InsertPrice(tx *sql.Tx, priceID uuid.UUID, flightID uuid.UUID, seatType enums.SeatType, price int) error {
	_, err := tx.Exec(`
		INSERT INTO pricing (price_id, flight_id, seat_type, price, created_on)
		VALUES ($1, $2, $3, $4, $5)
	`,
		priceID, flightID, seatType, price, time.Now(),
	)
	return err
}
