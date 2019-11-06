package database

import (
	"database/sql"

	"github.com/alecbcs/gatekey/token"
)

// Add checks if an entry is already in the database and
// if found updates the entry, else it adds the entry.
func Add(db *sql.DB, entry token.Token) (bool, error) {
	result, err := Get(db, entry.Value)
	if err != nil {
		return false, err
	}
	if result.Value != "" {
		return false, nil
	}
	err = Insert(db, entry)
	if err != nil {
		return false, err
	}
	return true, nil
}

// Insert adds a new entry into the database.
func Insert(db *sql.DB, entry token.Token) error {
	// Ping to check that database connection still exists.
	err := db.Ping()
	if err != nil {
		return err
	}

	stmt, err := db.Prepare(
		"INSERT INTO tokens(" +
			"token," +
			"machineID," +
			"jobID" +
			") VALUES(?,?,?);")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(
		entry.Value,
		entry.MachineID,
		entry.JobID)

	if err != nil {
		return err
	}
	return nil
}

// // Update patches an existing db entry with new data.
// func Update(db *sql.DB, entry *results.Entry) {
// 	err := db.Ping()
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	stmt, err := db.Prepare(
// 		"UPDATE appdata SET " +
// 			"latestURL = ?," +
// 			"latestVERSION = ?," +
// 			"currentURL = ?," +
// 			"currentVERSION = ?," +
// 			"upToDate = ?" +
// 			"WHERE id = ?;")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	_, err = stmt.Exec(
// 		entry.LatestURL,
// 		strings.Join(entry.LatestVersion, "."),
// 		entry.CurrentURL,
// 		strings.Join(entry.CurrentVersion, "."),
// 		entry.UpToDate,
// 		entry.ID)

// 	if err != nil {
// 		log.Fatal(err)
// 	}
// }
