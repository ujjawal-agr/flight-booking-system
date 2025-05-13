package booking

import (
	"database/sql"
	"errors"
	"flight-booking-system/internal/db"
	"fmt"
)

func BookFlightController(flightCode string, customerName string, customerContact string, seatInfo []SeatInfo) ([]Bill, error) {
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

	noOfSeats := len(seatInfo)

	bookingID, err := InsertBooking(tx, flightID, customerName, customerContact, noOfSeats)
	if err != nil {
		return nil, err
	}

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

func CancelBookingController(flightCode string, seats []Seat) error {
	conn := db.GetDB()
	tx, err := conn.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	flightID, err := GetFlightID(tx, flightCode)
	if err != nil {
		return err
	}
	for _, seat := range seats {
		seatType, seatNo := seat.SeatType, seat.SeatNo
		seatID, err := GetSeatID(tx, seatType, seatNo, flightID)
		if err != nil {
			return err
		}
		bookingID, err := GetBookingID(tx, seatID)
		if err != nil {
			return err
		}
		err = UpdateBookings(tx, bookingID)
		if err != nil {
			return err
		}
		err = UpdateSeats(tx, seatID)
		if err != nil {
			return err
		}
		err = UpdateFlights(tx, flightID)
		if err != nil {
			return err
		}
		err = UpdateBookingSeatMapping(tx, seatID)
		if err != nil {
			return err
		}
	}

	err = tx.Commit()
	if err != nil {
		return err
	}
	return nil
}
