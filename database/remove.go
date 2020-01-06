package database

import (
	"database/sql"
	"errors"

	"github.com/alecbcs/gatekey/token"
)

// Remove deletes a token entry from the database.
func Remove(db *sql.DB, entry token.Token) error {
	if entry.Value == "" {
		return errors.New("error removing empty value")
	}
	result, err := Get(db, entry.Value)
	if err != nil {
		return err
	}

	stmt, err := db.Prepare(
		"DELETE FROM tokens WHERE token = ?")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(result.Value)
	return err
}
