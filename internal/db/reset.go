package db

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"log"
)

func ResetDatabase(db *sql.DB, schemaPath string) error {
	// Drop and recreate the public schema
	_, err := db.Exec(`DROP SCHEMA public CASCADE; CREATE SCHEMA public;`)
	if err != nil {
		return fmt.Errorf("failed to reset schema: %v", err)
	}
	log.Println("Schema reset successfully.")

	// Read schema.sql file
	schemaBytes, err := ioutil.ReadFile(schemaPath)
	if err != nil {
		return fmt.Errorf("failed to read schema file: %v", err)
	}

	// Execute schema.sql contents
	_, err = db.Exec(string(schemaBytes))
	if err != nil {
		return fmt.Errorf("failed to execute schema: %v", err)
	}
	log.Println("Schema re-applied successfully.")

	return nil
}
