package handlers

import (
	"errors"
	"fmt"
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

	render.NoContent(w, r)
}

func (h *ItemsHandler) DisableItem(w http.ResponseWriter, r *http.Request) {
	itemID := chi.URLParam(r, "itemID")

	err := h.query.DisableItem(r.Context(), itemID)
	if err != nil {
		render.Render(w, r, ErrorServer(err))
		return
	}

	render.NoContent(w, r)
}

func (h *ItemsHandler) GetItems(w http.ResponseWriter, r *http.Request) {
	items, err := h.query.GetAllItems(r.Context())
	if err != nil {
		render.Render(w, r, ErrorServer(err))
		return
	}

	render.JSON(w, r, items)
}

func (h *ItemsHandler) GetItemsByCategory(w http.ResponseWriter, r *http.Request) {
	items, err := h.query.GetAllItems(r.Context())
	if err != nil {
		render.Render(w, r, ErrorServer(err))
		return
	}

	res := struct {
		Characters map[string]store.Item `json:"characters"`
		Summons    map[string]store.Item `json:"summons"`
		Weapons    map[string]store.Item `json:"weapons"`
	}{
		Characters: map[string]store.Item{},
		Summons:    map[string]store.Item{},
		Weapons:    map[string]store.Item{},
	}

	for _, item := range items {
		switch item.Category {
		case "character":
			res.Characters[item.ID] = item
		case "summon":
			res.Summons[item.ID] = item
		case "weapon":
			res.Weapons[item.ID] = item
		default:
			render.Render(w, r, ErrorServer(fmt.Errorf("Unknown category %v for item %v", item.Category, item.Name)))
			return
		}
	}

	render.JSON(w, r, res)
}

func (h *ItemsHandler) GetEnabledItems(w http.ResponseWriter, r *http.Request) {
	items, err := h.query.GetEnabledItems(r.Context())
	if err != nil {
		render.Render(w, r, ErrorServer(err))
		return
	}

	render.JSON(w, r, items)
}

type ItemValueRequest struct {
	Value string `json:"value"`
	Color string `json:"color"`
}

func (p *ItemValueRequest) Bind(r *http.Request) error {
	if p.Value == "" {
		return errors.New("Value field is required")
	}
	if p.Color == "" {
		return errors.New("Color field is required")
	}

	return nil
}

func (h *ItemsHandler) GetItemValues(w http.ResponseWriter, r *http.Request) {
	itemID := chi.URLParam(r, "itemID")

	values, err := h.query.GetItemValues(r.Context(), itemID)
	if err != nil {
		render.Render(w, r, ErrorInvalidRequest(err))
		return
	}

	if values == nil {
		values = []store.ItemValue{}
	}

	render.JSON(w, r, values)
}

func (h *ItemsHandler) CreateItemValue(w http.ResponseWriter, r *http.Request) {
	itemID := chi.URLParam(r, "itemID")

	var data ItemValueRequest
	if err := render.Bind(r, &data); err != nil {
		render.Render(w, r, ErrorInvalidRequest(err))
		return
	}

	dbValue, err := h.query.CreateItemValue(r.Context(), store.CreateItemValueParams{
		ItemID: itemID,
		Value:  data.Value,
		Color:  data.Color,
	})
	if err != nil {
		render.Render(w, r, ErrorInvalidRequest(err))
		return
	}

	render.JSON(w, r, dbValue)
}

func (h *ItemsHandler) DeleteItemValue(w http.ResponseWriter, r *http.Request) {
	itemID := chi.URLParam(r, "itemID")
	valueID := r.Context().Value(contextKey("valueID")).(int64)

	err := h.query.DeleteItemValue(r.Context(), store.DeleteItemValueParams{
		ItemID: itemID,
		ID:     valueID,
	})
	if err != nil {
		render.Render(w, r, ErrorInvalidRequest(err))
		return
	}

	render.NoContent(w, r)
}
