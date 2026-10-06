package handlers

import (
	"context"
	"database/sql"
	"errors"
	"net/http"
	"strconv"

	"github.com/Quak1/chuy-gbf/internal/store"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

type Middleware struct {
	query *store.Queries
}

func NewMiddlware(query *store.Queries) *Middleware {
	return &Middleware{
		query: query,
	}
}

type contextKey string

const (
	UserContextKey contextKey = "user"
)

func (m *Middleware) URLIntID(key string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			v, err := strconv.ParseInt(chi.URLParam(r, key), 10, 64)
			if err != nil {
				render.Render(w, r, ErrorInvalidRequest(err))
				return
			}

			ctx := context.WithValue(r.Context(), contextKey(key), v)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func (m *Middleware) UserCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userID := r.Context().Value(contextKey("userID")).(int64)

		user, err := m.query.GetUser(r.Context(), userID)
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				render.Render(w, r, ErrorNotFound)
				return
			}
			render.Render(w, r, ErrorServer(err))
			return
		}

		ctx := context.WithValue(r.Context(), UserContextKey, user)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
