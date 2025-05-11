package db

import (
	"database/sql"
	"log"
)

var dbInstance *sql.DB

func InitDB(connStr string) {
	var err error
	dbInstance, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Failed to open DB:", err)
	}

	if err = dbInstance.Ping(); err != nil {
		log.Fatal("Failed to connect to DB:", err)
	}
}

func GetDB() *sql.DB {
	return dbInstance
}
