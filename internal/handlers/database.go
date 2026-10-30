package handlers

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/Quak1/chuy-gbf/internal/store"
	"github.com/go-chi/render"
)

type DatabaseHandler struct {
	query *store.Queries
	db    *sql.DB
}

func NewDatabaseHandler(q *store.Queries, db *sql.DB) *DatabaseHandler {
	return &DatabaseHandler{
		query: q,
		db:    db,
	}
}

func (h *DatabaseHandler) BackupDatabase(w http.ResponseWriter, r *http.Request) {
	tmpDir, err := os.MkdirTemp("", "sqlite-backup-*")
	if err != nil {
		http.Error(w, "Failed to initialize backup storage", http.StatusInternalServerError)
		return
	}
	defer os.Remove(tmpDir)

	tmpPath := filepath.Join(tmpDir, "backup.db")

	if _, err = h.db.ExecContext(r.Context(), "VACUUM main INTO ?", tmpPath); err != nil {
		log.Printf("Database vacuum failed: %v", err)
		render.Render(w, r, ErrorServer(errors.New("Failed to snapshot database")))
		return
	}

	f, err := os.Open(tmpPath)
	if err != nil {
		log.Printf("Database backup failed: %v", err)
		render.Render(w, r, ErrorServer(errors.New("Failed to read completed backup")))
		return
	}
	defer f.Close()

	fileInfo, err := f.Stat()
	if err != nil {
		log.Printf("Database backup failed: %v", err)
		render.Render(w, r, ErrorServer(errors.New("Failed to calculate download payload size")))
		return
	}

	filename := fmt.Sprintf("backup-%s.db", time.Now().Format("20060102_150405"))
	w.Header().Set("Content-Type", "application/vnd.sqlite3")
	w.Header().Set("Content-Disposition", fmt.Sprintf(`attachment; filename="%s"`, filename))

	http.ServeContent(w, r, filename, fileInfo.ModTime(), f)
}
