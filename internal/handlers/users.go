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

func (h *UserHandler) ListUsers(w http.ResponseWriter, r *http.Request) {
	users, err := h.query.GetAllUsers(r.Context())
	if err != nil {
		render.Render(w, r, ErrorServer(err))
		return
	}

	render.JSON(w, r, users)
}

type UserRequest struct {
	Username string `json:"username"`
}

func (u *UserRequest) Bind(r *http.Request) error {
	if u.Username == "" {
		return errors.New("username field is required")
	}

	return nil
}

func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var user UserRequest

	if err := render.Bind(r, &user); err != nil {
		render.Render(w, r, ErrorInvalidRequest(err))
		return
	}

	if err := h.query.CreateUser(r.Context(), store.CreateUserParams{Username: user.Username}); err != nil {
		render.Render(w, r, ErrorServer(err))
		return
	}

	render.JSON(w, r, user)
}

func (h *UserHandler) GetUserItems(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value(UserContextKey).(store.User)

	userItems, err := h.query.GetUserItems(r.Context(), user.ID)
	if err != nil {
		render.Render(w, r, ErrorServer(err))
		return
	}

	render.JSON(w, r, userItems)
}

type UserItemRequest struct {
	Value string `json:"value"`
}

func (u *UserItemRequest) Bind(r *http.Request) error {
	if u.Value == "" {
		return errors.New("value field is required")
	}

	return nil
}

func (h *UserHandler) SetUserItem(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value(UserContextKey).(store.User)
	itemID := r.Context().Value(contextKey("itemID")).(int64)

	var value UserItemRequest
	if err := render.Bind(r, &value); err != nil {
		render.Render(w, r, ErrorInvalidRequest(err))
		return
	}

	err := h.query.SetUserItem(r.Context(), store.SetUserItemParams{
		UserID: user.ID,
		ItemID: itemID,
		Value:  value.Value,
	})
	if err != nil {
		render.Render(w, r, ErrorInvalidRequest(err))
		return
	}

	render.JSON(w, r, "OK")
}
