package handlers

import (
	"errors"
	"net/http"

	"github.com/Quak1/chuy-gbf/internal/store"
	"github.com/go-chi/chi/v5"
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

	dbUser, err := h.query.CreateUser(r.Context(), store.CreateUserParams{Username: user.Username})
	if err != nil {
		render.Render(w, r, ErrorServer(err))
		return
	}

	render.JSON(w, r, dbUser)
}

func (h *UserHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value(UserContextKey).(store.User)

	render.JSON(w, r, struct {
		ID       int64  `json:"id"`
		Username string `json:"username"`
		Comment  string `json:"comment"`
	}{
		ID:       user.ID,
		Username: user.Username,
		Comment:  user.Comment,
	})
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
	Value    string `json:"value"`
	Username string `json:"username"`
}

func (u *UserItemRequest) Bind(r *http.Request) error {
	if u.Value == "" {
		return errors.New("value field is required")
	}
	if u.Username == "" {
		return errors.New("username field is required")
	}

	return nil
}

func (h *UserHandler) SetUserItem(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value(UserContextKey).(store.User)
	itemID := chi.URLParam(r, "itemID")

	var data UserItemRequest
	if err := render.Bind(r, &data); err != nil {
		render.Render(w, r, ErrorInvalidRequest(err))
		return
	}

	if user.Username != data.Username {
		render.Render(w, r, ErrorForbidden)
		return
	}

	err := h.query.SetUserItem(r.Context(), store.SetUserItemParams{
		UserID: user.ID,
		ItemID: itemID,
		Value:  data.Value,
	})
	if err != nil {
		render.Render(w, r, ErrorInvalidRequest(err))
		return
	}

	render.JSON(w, r, "OK")
}

type UserCommentRequest struct {
	Comment  string `json:"comment"`
	Username string `json:"username"`
}

func (u *UserCommentRequest) Bind(r *http.Request) error {
	if u.Username == "" {
		return errors.New("username field is required")
	}

	return nil
}

func (h *UserHandler) SetUserComment(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value(UserContextKey).(store.User)

	var data UserCommentRequest
	if err := render.Bind(r, &data); err != nil {
		render.Render(w, r, ErrorInvalidRequest(err))
		return
	}

	if user.Username != data.Username {
		render.Render(w, r, ErrorForbidden)
		return
	}

	err := h.query.SetUserComment(r.Context(), store.SetUserCommentParams{
		Comment: data.Comment,
		ID:      user.ID,
	})
	if err != nil {
		render.Render(w, r, ErrorInvalidRequest(err))
		return
	}

	render.JSON(w, r, "OK")
}
