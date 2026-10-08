package database

import (
	"cmp"
	"database/sql"
	"fmt"
	"os"

	_ "modernc.org/sqlite"
)

func InitDB() (*sql.DB, error) {
	dbPath := cmp.Or(os.Getenv("DB_PATH"), "data.db")
	dsn := fmt.Sprintf("file:%s?", dbPath) +
		"_pragma=foreign_keys(1)&" +
		"_pragma=journal_mode(WAL)&" +
		"_pragma=busy_timeout(5000)"

	db, err := sql.Open("sqlite", dsn)
	if err != nil {
		return nil, fmt.Errorf("Failed to open database: %w", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("Failed to ping database: %w", err)
	}

	return db, nil
}
