package booking

import (
	"database/sql"
	"errors"
	"flight-booking-system/internal/db"
	"flight-booking-system/internal/enums"
	"fmt"
	"github.com/google/uuid"
)

type Bill struct {
	PassengerName string         `json:"passenger_name"`
	FlightCode    string         `json:"flight_code"`
	SeatType      enums.SeatType `json:"seat_type"`
	SeatNo        int            `json:"seat_no"`
	Amount        int            `json:"amount"`
}

func BookFlightController(flightCode string, customerName string, customerContact string, seatInfo []Seat) ([]Bill, error) {
	conn := db.GetDB()
	tx, err := conn.Begin()
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	flightID, err := GetFlightID(tx, flightCode)
	if err != nil {
		return nil, err
	}

	// Step 1: Insert booking
	bookingID, err := InsertBooking(tx, flightID, customerName, customerContact)
	if err != nil {
		return nil, err
	}

	// Step 3: Process seats
	var bills []Bill
	for _, seat := range seatInfo {
		seatID, seatNo, err := GetAvailableSeat(tx, flightID, seat.SeatType)
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("no available seat of type %s", seat.SeatType)
		} else if err != nil {
			return nil, err
		}

		err = BookSeat(tx, seatID)
		if err != nil {
			return nil, err
		}

		err = InsertBookingSeatMapping(tx, bookingID, seatID, seat.Name, seat.Age, seat.Gender)
		if err != nil {
			return nil, err
		}

		price, err := GetSeatPrice(tx, flightID, seat.SeatType)
		if err != nil {
			return nil, err
		}

		bills = append(bills, Bill{
			PassengerName: seat.Name,
			FlightCode:    flightCode,
			SeatType:      seat.SeatType,
			SeatNo:        seatNo,
			Amount:        price,
		})
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	return bills, nil
}

func CancelBookingController(bookingID uuid.UUID) error {
	return CancelBooking(bookingID)
}
