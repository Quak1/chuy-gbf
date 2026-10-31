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
