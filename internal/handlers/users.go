package handlers

import (
	"errors"
	"net/http"

	"github.com/Quak1/chuy-gbf/internal/store"
	"github.com/go-chi/render"
)

type UserHandler struct {
	query *store.Queries
}

func NewUsersHandler(query *store.Queries) *UserHandler {
	return &UserHandler{
		query: query,
	}
}

type Data struct {
	Data string `json:"data"`
}

func (h *UserHandler) Ping(w http.ResponseWriter, r *http.Request) {
	d := struct {
		Data string `json:"data"`
	}{
		Data: "pong",
	}
	render.JSON(w, r, d)
}

func (h *UserHandler) ListUsers(w http.ResponseWriter, r *http.Request) {
	users, err := h.query.GetAllUsers(r.Context())
	if err != nil {
		render.Render(w, r, ErrorServer(err))
	}

	render.JSON(w, r, users)
}

type UserRequest struct {
	Name string `json:"name"`
}

func (u *UserRequest) Bind(r *http.Request) error {
	if u.Name == "" {
		return errors.New("name is required")
	}

	return nil
}

func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var user UserRequest

	if err := render.Bind(r, &user); err != nil {
		render.Render(w, r, ErrorInvalidRequest(err))
		return
	}

	if err := h.query.CreateUser(r.Context(), store.CreateUserParams{Name: user.Name}); err != nil {
		render.Render(w, r, ErrorServer(err))
		return
	}

	render.JSON(w, r, user)
}
