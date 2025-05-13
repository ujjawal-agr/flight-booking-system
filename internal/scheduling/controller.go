package scheduling

import (
	"errors"
	"flight-booking-system/internal/booking"
	"flight-booking-system/internal/db"
	"fmt"
	"github.com/google/uuid"
	"time"
)

func ScheduleFlightController(req ScheduleFlightRequest) error {
	conn := db.GetDB()
	tx, err := conn.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	// 1. Basic validation
	if req.FlightCode == "" || req.Source == "" || req.Destination == "" || req.Company == "" {
		return errors.New("missing required fields")
	}
	if req.Date.Before(time.Now()) {
		return errors.New("flight date must be in the future")
	}

	// 2. Create flight ID
	flightID := uuid.New()

	// 3. Insert flight
	err = InsertFlight(tx, flightID, req)
	if err != nil {
		return fmt.Errorf("failed to insert flight: %w", err)
	}

	// 4. Insert seat records for each seat type

	for seatType, count := range req.SeatCounts {
		for i := 1; i <= count; i++ {
			seatID := uuid.New()
			//seatNo := fmt.Sprintf("%s-%03d", seatType[:1], i) // e.g., W-001
			seatNo := i

			err := InsertSeat(tx, seatID, flightID, seatNo, seatType)
			if err != nil {
				return fmt.Errorf("failed to insert seat: %w", err)
			}
		}
	}

	// 5. Insert pricing
	for seatType, price := range req.Prices {
		priceID := uuid.New()
		err := InsertPrice(tx, priceID, flightID, seatType, price)
		if err != nil {
			return fmt.Errorf("failed to insert price: %w", err)
		}
	}

	err = tx.Commit()
	if err != nil {
		return err
	}
	return nil
}

func CancelFlightController(req CancelFlightRequest) error {
	conn := db.GetDB()
	tx, err := conn.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	flightCode := req.FlightCode

	flightID, err := booking.GetFlightID(tx, flightCode)
	if err != nil {
		return err
	}

	err = UpdateFlightStatus(tx, flightID)
	if err != nil {
		return err
	}

	bookedSeats := GetBookedSeats(tx, flightID)

	err = UpdateSeats(tx, flightID)
	if err != nil {
		return err
	}

	err = UpdatePricing(tx, flightID)
	if err != nil {
		return err
	}

	err = UpdateBookings(tx, flightID)
	if err != nil {
		return err
	}

	if len(bookedSeats) > 0 {
		for _, seat := range bookedSeats {
			err = UpdateBookingSeatMapping(tx, seat)
			if err != nil {
				return err
			}
		}
	}

	err = tx.Commit()
	if err != nil {
		return err
	}
	return nil
}
