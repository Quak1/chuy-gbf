package handlers

import (
	"database/sql"
	"errors"
	"net/http"

	"github.com/Quak1/chuy-gbf/internal/store"
	"github.com/go-chi/render"
)

type TeamsHandler struct {
	query *store.Queries
	db    *sql.DB
}

func NewTeamsHandler(q *store.Queries, db *sql.DB) *TeamsHandler {
	return &TeamsHandler{
		query: q,
		db:    db,
	}
}

type TeamRequest struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

func (p *TeamRequest) Bind(r *http.Request) error {
	if p.Name == "" {
		return errors.New("name field is required")
	}

	return nil
}

func (h *TeamsHandler) ListTeams(w http.ResponseWriter, r *http.Request) {
	teams, err := h.query.ListTeams(r.Context())
	if err != nil {
		render.Render(w, r, ErrorServer(err))
		return
	}

	render.JSON(w, r, teams)
}

type SlotItem struct {
	store.Item
	Position int64 `json:"position"`
}

func (h *TeamsHandler) GetTeam(w http.ResponseWriter, r *http.Request) {
	teamID := r.Context().Value(TeamIDKey).(int64)

	dbTeam, err := h.query.GetTeam(r.Context(), teamID)
	if err != nil {
		HandleDatabaseError(w, r, err)
		return
	}

	teamSlots, err := h.query.GetTeamSlots(r.Context(), teamID)
	if err != nil {
		HandleDatabaseError(w, r, err)
		return
	}

	team := struct {
		Name       string     `json:"name"`
		URL        string     `json:"url"`
		ID         int64      `json:"id"`
		Characters []SlotItem `json:"characters"`
		Summons    []SlotItem `json:"summons"`
		Weapons    []SlotItem `json:"weapons"`
	}{
		ID:   dbTeam.ID,
		Name: dbTeam.Name,
		URL:  dbTeam.Url,
	}

	for _, i := range teamSlots {
		item := SlotItem{
			Item: store.Item{
				ID:       i.ID,
				Name:     i.Name,
				Element:  i.Element,
				Type:     i.Type,
				Series:   i.Series,
				Enabled:  i.Enabled,
				Category: i.Category,
				WikiPath: i.WikiPath,
			},
			Position: i.Position,
		}

		switch i.Category {
		case "character":
			team.Characters = append(team.Characters, item)
		case "summon":
			team.Summons = append(team.Summons, item)
		case "weapon":
			team.Weapons = append(team.Weapons, item)
		}
	}

	render.JSON(w, r, team)
}

func (h *TeamsHandler) CreateTeam(w http.ResponseWriter, r *http.Request) {
	var data TeamRequest
	if err := render.Bind(r, &data); err != nil {
		render.Render(w, r, ErrorInvalidRequest(err))
		return
	}

	teamId, err := h.query.CreateTeam(r.Context(), data.Name)
	if err != nil {
		render.Render(w, r, ErrorServer(err))
		return
	}

	res := map[string]int64{}
	res["team_id"] = teamId

	render.JSON(w, r, res)
}

func (h *TeamsHandler) UpdateTeam(w http.ResponseWriter, r *http.Request) {
	teamID := r.Context().Value(TeamIDKey).(int64)

	var data TeamRequest
	if err := render.Decode(r, &data); err != nil {
		render.Render(w, r, ErrorInvalidRequest(err))
		return
	}

	tx, err := h.db.Begin()
	if err != nil {
		render.Render(w, r, ErrorServer(err))
		return
	}
	defer tx.Rollback()

	qtx := h.query.WithTx(tx)

	err = qtx.UpdateTeam(r.Context(), store.UpdateTeamParams{
		Name: data.Name,
		Url:  data.URL,
		ID:   teamID,
	})
	if err != nil {
		render.Render(w, r, ErrorServer(err))
		return
	}

	err = qtx.SetTeamUpdated(r.Context(), teamID)
	if err != nil {
		render.Render(w, r, ErrorServer(err))
		return
	}

	err = tx.Commit()
	if err != nil {
		render.Render(w, r, ErrorServer(err))
		return
	}

	render.NoContent(w, r)
}

func (h *TeamsHandler) DeleteTeam(w http.ResponseWriter, r *http.Request) {
	teamID := r.Context().Value(TeamIDKey).(int64)

	err := h.query.DeleteTeam(r.Context(), teamID)
	if err != nil {
		render.Render(w, r, ErrorServer(err))
		return
	}

	render.NoContent(w, r)
}

type TeamSlotRequest struct {
	Category string `json:"category"`
	Position int64  `json:"position"`
	ItemId   string `json:"item_id"`
}

func (t *TeamSlotRequest) Bind(r *http.Request) error {
	if t.Category == "" {
		return errors.New("category field is required")
	}

	if t.Position == 0 {
		return errors.New("position field is required")
	}

	return nil
}

func (h *TeamsHandler) SetTeamSlot(w http.ResponseWriter, r *http.Request) {
	teamID := r.Context().Value(TeamIDKey).(int64)

	var data TeamSlotRequest
	if err := render.Bind(r, &data); err != nil {
		render.Render(w, r, ErrorInvalidRequest(err))
		return
	}

	if data.ItemId == "" {
		render.Render(w, r, ErrorInvalidRequest(errors.New("item_id field is required")))
		return
	}

	tx, err := h.db.Begin()
	if err != nil {
		render.Render(w, r, ErrorServer(err))
		return
	}
	defer tx.Rollback()

	qtx := h.query.WithTx(tx)

	err = qtx.SetSlot(r.Context(), store.SetSlotParams{
		TeamID:   teamID,
		Category: data.Category,
		Position: data.Position,
		ItemID:   data.ItemId,
	})
	if err != nil {
		HandleDatabaseError(w, r, err)
		return
	}

	if err := qtx.SetTeamUpdated(r.Context(), teamID); err != nil {
		HandleDatabaseError(w, r, err)
		return
	}

	err = tx.Commit()
	if err != nil {
		render.Render(w, r, ErrorServer(err))
		return
	}

	render.NoContent(w, r)
}

func (h *TeamsHandler) ClearTeamSlot(w http.ResponseWriter, r *http.Request) {
	teamID := r.Context().Value(TeamIDKey).(int64)

	var data TeamSlotRequest
	if err := render.Bind(r, &data); err != nil {
		render.Render(w, r, ErrorInvalidRequest(err))
		return
	}

	tx, err := h.db.Begin()
	if err != nil {
		render.Render(w, r, ErrorServer(err))
		return
	}
	defer tx.Rollback()

	qtx := h.query.WithTx(tx)

	err = qtx.ClearSlot(r.Context(), store.ClearSlotParams{
		TeamID:   teamID,
		Category: data.Category,
		Position: data.Position,
	})
	if err != nil {
		HandleDatabaseError(w, r, err)
		return
	}

	if err := qtx.SetTeamUpdated(r.Context(), teamID); err != nil {
		HandleDatabaseError(w, r, err)
		return
	}

	err = tx.Commit()
	if err != nil {
		render.Render(w, r, ErrorServer(err))
		return
	}

	render.NoContent(w, r)
}
