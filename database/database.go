package database

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3" // Import sqlite3 driver for database interaction.
)

// Open opens a database and creates one if not found.
func Open(databaseName string) (db *sql.DB) {
	var err error
	db, err = sql.Open("sqlite3", databaseName)
	if err != nil {
		log.Fatal(err)
	}
	// Create the appdata table if is doesn't already exist.
	// This will also create the database if it doesn't exist.
	_, err = db.Exec(
		"CREATE TABLE IF NOT EXISTS tokens(" +
			"token TEXT PRIMARY KEY," +
			"machineID TEXT," +
			"jobID TEXT" +
			");")
	if err != nil {
		log.Fatal(err)
	}
	return db
}
