package routes

import (
	"database/sql"

	"github.com/Quak1/chuy-gbf/internal/handlers"
	"github.com/Quak1/chuy-gbf/internal/store"
	"github.com/go-chi/chi/v5"
)

func SetupRouter(db *sql.DB) *chi.Mux {
	q := store.New(db)
	r := chi.NewRouter()

	itemsHandler := handlers.NewItemsHandler(q)

	r.Get("/ping", itemsHandler.Ping)

	return r
}
