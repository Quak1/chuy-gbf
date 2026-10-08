package database

import (
	"cmp"
	"database/sql"
	"embed"
	"fmt"
	"os"

	"github.com/pressly/goose/v3"
	_ "modernc.org/sqlite"
)

//go:embed migrations/*.sql
var embedMigrations embed.FS

const dialect = "sqlite"

func InitDB() (*sql.DB, error) {
	dbPath := cmp.Or(os.Getenv("DB_PATH"), "data.db")
	dsn := fmt.Sprintf("file:%s?", dbPath) +
		"_foreign_keys=1&" +
		"_journal_mode=WAL&" +
		"_busy_timeout=5000"

	db, err := sql.Open(dialect, dsn)
	if err != nil {
		return nil, fmt.Errorf("Failed to open database: %w", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("Failed to ping database: %w", err)
	}

	// Run migrations
	goose.SetBaseFS(embedMigrations)

	if err := goose.SetDialect(dialect); err != nil {
		return nil, err
	}

	if err := goose.Up(db, "migrations"); err != nil {
		return nil, err
	}

	return db, nil
}
