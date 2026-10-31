-- name: CreateTeam :one
INSERT INTO teams (name, url) VALUES (?, "") RETURNING id;

-- name: UpdateTeam :exec
UPDATE teams
SET name = ?, url = ?, updated_at = datetime('now')
WHERE id = ?;

-- name: DeleteTeam :exec
DELETE FROM teams WHERE id = ?;

-- name: ListTeams :many
SELECT * FROM teams ORDER BY id;

-- name: SetTeamUpdated :exec
UPDATE teams SET updated_at = datetime('now') WHERE id = ?;

-- name: GetTeam :one
SELECT * FROM teams
WHERE teams.id = ?;

-- name: GetTeamSlots :many
SELECT
  ts.position,
  i.*
FROM teams t
JOIN team_slots ts ON t.id = ts.team_id
JOIN items i ON ts.item_id = i.id
WHERE t.id = ?
ORDER BY
  ts.category,
  ts.position;

-- name: ListTeamsByItemId :many
SELECT t.* FROM teams t
JOIN team_slots ts ON ts.team_id = t.id
WHERE ts.item_id = ?;

---------------------------
--         SLOTS         --
---------------------------
-- name: SetSlot :exec
INSERT INTO team_slots (team_id, category, position, item_id)
VALUES (?, ?, ?, ?)
ON CONFLICT (team_id, category, position)
DO UPDATE SET item_id = excluded.item_id;

-- name: ClearSlot :exec
DELETE FROM team_slots
WHERE team_id = ? AND category = ? AND position = ?;
