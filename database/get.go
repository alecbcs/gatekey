package database

import (
	"database/sql"

	"github.com/alecbcs/gatekey/token"
)

// Get finds and returns an entry from the database.
func Get(db *sql.DB, value string) (token.Token, error) {
	var (
		result token.Token
	)

	err := db.Ping()
	if err != nil {
		return result, err
	}
	row, err := db.Query("SELECT * FROM tokens WHERE token = ?", value)
	if err != nil {
		return result, err
	}
	defer row.Close()
	if !row.Next() {
		return result, nil
	}
	err = row.Scan(
		&result.Value,
		&result.MachineID,
		&result.JobID)
	if err != nil {
		return result, err
	}
	return result, nil
}
