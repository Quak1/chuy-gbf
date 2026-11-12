package handlers

import (
	"context"
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

type SlotItem struct {
	store.Item
	Position int64 `json:"position"`
}

type TeamWeapon struct {
	store.GetTeamSlotsRow
	Skills [3 + 1]*store.GetTeamWeaponsSkillsRow `json:"skills"`
}

type Team struct {
	Name       string                        `json:"name"`
	URL        string                        `json:"url"`
	ID         int64                         `json:"id"`
	Characters [6 + 1]*store.GetTeamSlotsRow `json:"character"`
	Summons    [7 + 1]*store.GetTeamSlotsRow `json:"summon"`
	Weapons    [10 + 1]*TeamWeapon           `json:"weapon"`
}

func (p *TeamRequest) Bind(r *http.Request) error {
	if p.Name == "" {
		return errors.New("name field is required")
	}

	return nil
}

func (h *TeamsHandler) ListTeams(w http.ResponseWriter, r *http.Request) {
	dbTeams, err := h.query.ListTeams(r.Context())
	if err != nil {
		render.Render(w, r, ErrorServer(err))
		return
	}

	// TODO handle concurrency style
	var teams []*Team
	for _, t := range dbTeams {
		team, err := h.processTeam(r.Context(), t.ID)
		if err != nil {
			HandleDatabaseError(w, r, err)
			return
		}

		teams = append(teams, team)
	}

	render.JSON(w, r, teams)
}

func (h *TeamsHandler) GetTeam(w http.ResponseWriter, r *http.Request) {
	teamID := r.Context().Value(TeamIDKey).(int64)

	team, err := h.processTeam(r.Context(), teamID)
	if err != nil {
		HandleDatabaseError(w, r, err)
		return
	}

	render.JSON(w, r, team)
}

func (h *TeamsHandler) processTeam(ctx context.Context, teamID int64) (*Team, error) {
	dbTeam, err := h.query.GetTeam(ctx, teamID)
	if err != nil {
		return nil, err
	}

	teamSlots, err := h.query.GetTeamSlots(ctx, teamID)
	if err != nil {
		return nil, err
	}

	team := Team{
		ID:   dbTeam.ID,
		Name: dbTeam.Name,
		URL:  dbTeam.Url,
	}

	for _, item := range teamSlots {
		switch item.Category {
		case "character":
			team.Characters[item.Position] = &item
		case "summon":
			team.Summons[item.Position] = &item
		case "weapon":
			w := &TeamWeapon{GetTeamSlotsRow: item}

			dbSkills, err := h.query.GetTeamWeaponsSkills(ctx, team.ID)
			if err != nil {
				return nil, err
			}

			var skills [4]*store.GetTeamWeaponsSkillsRow
			for _, s := range dbSkills {
				skills[s.Position] = &s
			}
			w.Skills = skills

			team.Weapons[item.Position] = w
		}
	}

	return &team, nil
}

func (h *TeamsHandler) CreateTeam(w http.ResponseWriter, r *http.Request) {
	teamId, err := h.query.CreateTeam(r.Context(), "New Team")
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

	err = qtx.ClearSlotWeaponSkills(r.Context(), store.ClearSlotWeaponSkillsParams{
		TeamID:       teamID,
		SlotCategory: data.Category,
		SlotPosition: data.Position,
	})
	if err != nil {
		HandleDatabaseError(w, r, err)
		return
	}

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

type TeamWeaponSkillRequest struct {
	SkillId int64 `json:"skill_id"`
}

func (t *TeamWeaponSkillRequest) Bind(r *http.Request) error {
	if t.SkillId == 0 {
		return errors.New("skill_id field is required")
	}

	return nil
}

func (h *TeamsHandler) SetWeaponSkill(w http.ResponseWriter, r *http.Request) {
	teamID := r.Context().Value(TeamIDKey).(int64)
	slotPosition := r.Context().Value(SlotPositionKey).(int64)
	skillPosition := r.Context().Value(SkillPositionKey).(int64)

	var data TeamWeaponSkillRequest
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

	slot, err := qtx.GetSlot(r.Context(), store.GetSlotParams{
		TeamID:   teamID,
		Category: "weapon",
		Position: slotPosition,
	})
	if err != nil {
		HandleDatabaseError(w, r, err)
		return
	}

	err = qtx.SetWeaponSkill(r.Context(), store.SetWeaponSkillParams{
		WeaponSkillID: data.SkillId,
		Position:      skillPosition,
		TeamID:        slot.TeamID,
		ItemID:        slot.ItemID,
		SlotCategory:  slot.Category,
		SlotPosition:  slot.Position,
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

func (h *TeamsHandler) ClearWeaponSkill(w http.ResponseWriter, r *http.Request) {
	teamID := r.Context().Value(TeamIDKey).(int64)
	slotPosition := r.Context().Value(SlotPositionKey).(int64)
	skillPosition := r.Context().Value(SkillPositionKey).(int64)

	var data TeamWeaponSkillRequest
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

	err = qtx.ClearWeaponSkill(r.Context(), store.ClearWeaponSkillParams{
		TeamID:       teamID,
		SlotPosition: slotPosition,
		SlotCategory: "weapon",
		Position:     skillPosition,
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
