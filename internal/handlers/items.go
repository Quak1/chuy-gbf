package handlers

import (
	"net/http"

	"github.com/Quak1/chuy-gbf/internal/store"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

type ItemsHandler struct {
	query *store.Queries
}

func NewItemsHandler(q *store.Queries) *ItemsHandler {
	return &ItemsHandler{
		query: q,
	}
}

func (h *ItemsHandler) EnableItem(w http.ResponseWriter, r *http.Request) {
	itemID := chi.URLParam(r, "itemID")

	err := h.query.EnableItem(r.Context(), itemID)
	if err != nil {
		render.Render(w, r, ErrorServer(err))
		return
	}

	render.JSON(w, r, "OK")
}

func (h *ItemsHandler) DisableItem(w http.ResponseWriter, r *http.Request) {
	itemID := chi.URLParam(r, "itemID")

	err := h.query.DisableItem(r.Context(), itemID)
	if err != nil {
		render.Render(w, r, ErrorServer(err))
		return
	}

	render.JSON(w, r, "OK")
}

func (h *ItemsHandler) GetItems(w http.ResponseWriter, r *http.Request) {
	items, err := h.query.GetAllItems(r.Context())
	if err != nil {
		render.Render(w, r, ErrorServer(err))
		return
	}

	render.JSON(w, r, items)
}

func (h *ItemsHandler) GetEnabledItems(w http.ResponseWriter, r *http.Request) {
	items, err := h.query.GetEnabledItems(r.Context())
	if err != nil {
		render.Render(w, r, ErrorServer(err))
		return
	}

	render.JSON(w, r, items)
}
