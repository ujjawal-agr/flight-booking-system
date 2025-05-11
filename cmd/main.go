package main

import (
	"database/sql"
	"flight-booking-system/internal/booking"
	"flight-booking-system/internal/db"
	"flight-booking-system/internal/scheduling"
	"flight-booking-system/internal/searching"
	"fmt"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"os"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Build connection string from env variables
	connStr := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)

	db.InitDB(connStr)

	DB := db.GetDB()

	defer func(DB *sql.DB) {
		err := DB.Close()
		if err != nil {
			log.Fatal("Failed to close DB:", err)
		}
	}(DB)

	err = db.ResetDatabase(DB, "internal/db/schema.sql")
	if err != nil {
		log.Fatalf("Database reset failed: %v", err)
	}

	fmt.Println("Connected to PostgreSQL successfully!")

	mux := http.NewServeMux()

	// Register booking-related routes
	booking.RegisterRoutes(mux)
	scheduling.RegisterRoutes(mux)
	searching.RegisterRoutes(mux)

	// Start the HTTP server
	log.Println("Server is running on http://localhost:8080")
	err = http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatalf("Server failed: %v", err)
	}

}
