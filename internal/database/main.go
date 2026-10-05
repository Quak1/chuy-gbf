package database

import (
	"database/sql"
	"fmt"

	_ "modernc.org/sqlite"
)

const (
	dbPath = "data.db"
)

func InitDB() (*sql.DB, error) {
	db, err := sql.Open("sqlite", dbPath)
	if err != nil {
		return nil, fmt.Errorf("Failed to open database: %w", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("Failed to ping database: %w", err)
	}

	return db, nil
}
