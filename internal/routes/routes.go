package routes

import (
	"database/sql"
	"io"
	"io/fs"
	"net/http"
	"path"
	"strings"

	"github.com/Quak1/chuy-gbf/internal/handlers"
	"github.com/Quak1/chuy-gbf/internal/store"
	"github.com/go-chi/chi/v5"
)

func SetupRouter(db *sql.DB, distFS fs.FS) *chi.Mux {
	q := store.New(db)
	r := chi.NewRouter()

	middleware := handlers.NewMiddlware(q)
	usersHandler := handlers.NewUsersHandler(q)
	itemsHandler := handlers.NewItemsHandler(q)

	r.Route("/api", func(r chi.Router) {
		r.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("Pong"))
		})

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

	r.Get("/*", func(w http.ResponseWriter, req *http.Request) {
		requestPath := path.Clean(req.URL.Path)

		file, err := distFS.Open(strings.TrimPrefix(requestPath, "/"))
		if err == nil {
			file.Close()
			http.FileServerFS(distFS).ServeHTTP(w, req)
			return
		}

		indexFile, err := distFS.Open("index.html")
		if err != nil {
			http.Error(w, "Index file not found", http.StatusInternalServerError)
			return
		}
		defer indexFile.Close()

		// Read and serve the index.html content
		stat, _ := indexFile.Stat()
		http.ServeContent(w, req, "index.html", stat.ModTime(), indexFile.(io.ReadSeeker))
	})

	return r
}
