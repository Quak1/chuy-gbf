package handlers

import (
	"net/http"

	"github.com/Quak1/chuy-gbf/internal/store"
	"github.com/go-chi/render"
)

type UserHandler struct {
	query *store.Queries
}

func NewItemsHandler(query *store.Queries) *UserHandler {
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
