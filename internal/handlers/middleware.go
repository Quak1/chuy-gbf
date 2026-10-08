package handlers

import (
	"context"
	"database/sql"
	"errors"
	"log"
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
	UsernameKey    contextKey = "username"
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

func (m *Middleware) IsAdminRole(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		username := r.Context().Value(UsernameKey).(string)

		user, err := m.query.GetUserByUsername(r.Context(), username)
		if err != nil {
			log.Println(err)
			render.Render(w, r, ErrorForbidden)
			return
		}

		if user.Role != "admin" {
			render.Render(w, r, ErrorForbidden)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func (m *Middleware) RequireUsername(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("username")
		if err != nil {
			if errors.Is(err, http.ErrNoCookie) {
				render.Render(w, r, ErrorInvalidRequest(err))
			}
			render.Render(w, r, ErrorServer(err))
		}

		ctx := context.WithValue(r.Context(), contextKey(UsernameKey), cookie.Value)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
