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
	middleware := handlers.NewMiddlware(q)

	r.Route("/users", func(r chi.Router) {
		r.Get("/", usersHandler.ListUsers)
		r.Post("/", usersHandler.CreateUser)

		r.Route("/{userID:[0-9]+}", func(r chi.Router) {
			r.Use(middleware.URLIntID("userID"))
			r.Use(middleware.UserCtx)

			r.Route("/items", func(r chi.Router) {
				r.Get("/", usersHandler.GetUserItems)
				r.With(middleware.URLIntID("itemID")).Post("/{itemID:[0-9]+}", usersHandler.SetUserItem)
			})
		})
	})

	return r
}
