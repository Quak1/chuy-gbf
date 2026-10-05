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

	usersHandler := handlers.NewUsersHandler(q)

	r.Get("/ping", usersHandler.Ping)
	r.Route("/users", func(r chi.Router) {
		r.Get("/", usersHandler.ListUsers)
		r.Post("/", usersHandler.CreateUser)
	})

	return r
}
