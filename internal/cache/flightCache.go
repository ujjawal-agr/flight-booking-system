package cache

import (
	"fmt"
	"github.com/redis/go-redis/v9"
	"time"
)

// SetFlightAsFullyBooked caches the flight key with an expiration
func SetFlightAsFullyBooked(flightCode string, date string) error {
	key := fmt.Sprintf("fullyBooked:%s:%s", flightCode, date)
	return Rdb.Set(Ctx, key, "1", 24*time.Hour).Err()
}

// IsFlightFullyBooked checks if the flight is cached
func IsFlightFullyBooked(flightCode string, date string) (bool, error) {
	key := fmt.Sprintf("fullyBooked:%s:%s", flightCode, date)
	val, err := Rdb.Get(Ctx, key).Result()
	if err == redis.Nil {
		return false, nil
	}
	return val == "1", err
}
