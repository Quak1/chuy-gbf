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

	middleware := handlers.NewMiddlware(q)
	usersHandler := handlers.NewUsersHandler(q)
	itemsHandler := handlers.NewItemsHandler(q)

	r.Route("/api", func(r chi.Router) {
		r.Route("/users", func(r chi.Router) {
			r.Get("/", usersHandler.ListUsers)
			r.Post("/", usersHandler.CreateUser)
			r.Get("/items", usersHandler.GetAllUsersItems)

			r.Route("/{userID:[0-9]+}", func(r chi.Router) {
				r.Use(middleware.URLIntID("userID"))
				r.Use(middleware.UserCtx)

				r.Get("/", usersHandler.GetUser)
				r.Post("/", usersHandler.SetUserComment)

				r.Route("/items", func(r chi.Router) {
					r.Get("/", usersHandler.GetUserItems)
					r.Post("/{itemID:[0-9]+}", usersHandler.SetUserItem)
				})
			})
		})

		r.Route("/items", func(r chi.Router) {
			r.Get("/", itemsHandler.GetItems)
			r.Get("/enabled", itemsHandler.GetEnabledItems)

			r.Route("/{itemID:[0-9]+}", func(r chi.Router) {
				r.Use(middleware.IsAdminRole)
				r.Post("/enable", itemsHandler.EnableItem)
				r.Post("/disable", itemsHandler.DisableItem)
			})
		})
	})

	return r
}
