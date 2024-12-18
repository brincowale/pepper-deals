package storage

import (
	"database/sql"
	"log"

	_ "modernc.org/sqlite"
)

var DB *sql.DB

type Deal struct {
	ID string
}

func init() {
	var err error
	DB, err = sql.Open("sqlite", "deals.db")
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	// Create the deals table if it doesn't exist
	_, err = DB.Exec(`CREATE TABLE IF NOT EXISTS deals (id TEXT PRIMARY KEY)`)
	if err != nil {
		log.Fatalf("failed to create table: %v", err)
	}
}

// InsertDeal inserts a new deal into the database
func InsertDeal(dealId string) bool {
	_, err := DB.Exec("INSERT INTO deals (id) VALUES (?)", dealId)
	if err != nil {
		log.Printf("failed to insert deal: %v", err)
		return false
	}
	return true
}
