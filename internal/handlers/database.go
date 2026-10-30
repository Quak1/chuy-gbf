package handlers

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/Quak1/chuy-gbf/internal/store"
	"github.com/Quak1/chuy-gbf/internal/updater"
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

func (h *DatabaseHandler) UpdateDBData(w http.ResponseWriter, r *http.Request) {
	var data struct {
		Force bool `form:"force"`
	}
	if err := render.Decode(r, &data); err != nil {
		render.Render(w, r, ErrorInvalidRequest(err))
		return
	}

	res := map[string]any{}

	if !data.Force {
		update, err := h.query.GetLastDBDataUpdate(r.Context())
		if err != nil && !errors.Is(err, sql.ErrNoRows) {
			render.Render(w, r, ErrorServer(err))
			return
		}

		if err == nil {
			c, err := updater.GetChangelog()
			if err != nil {
				render.Render(w, r, ErrorServer(err))
				return
			}

			t := time.UnixMilli(c.Timestamp)
			if update.CreatedAt.After(t) {
				res["message"] = fmt.Sprintf("Database already has latest data. Last DB Update: %v. Lastest data: %v", update.CreatedAt, t)
				render.JSON(w, r, res)
				return
			}
		}
	}

	err, errorMsgs := updater.LoadData(true, h.query)
	if err != nil {
		render.Render(w, r, ErrorServer(err))
		return
	}

	err = h.query.AddDBDataUpdate(context.Background(), store.AddDBDataUpdateParams{
		IsOk:     err == nil,
		Comments: strings.Join(errorMsgs, "\n\n"),
	})

	if errorMsgs != nil {
		res["message"] = "Completed with errors"
		res["errors"] = errorMsgs
		render.JSON(w, r, res)
		return
	}

	render.NoContent(w, r)
}

func (h *DatabaseHandler) GetLastUpdate(w http.ResponseWriter, r *http.Request) {
	update, err := h.query.GetLastDBDataUpdate(r.Context())
	if err != nil {
		render.Render(w, r, ErrorServer(err))
		return
	}

	render.JSON(w, r, update)
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
