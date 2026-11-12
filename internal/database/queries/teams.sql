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
-- name: GetSlot :one
SELECT * FROM team_slots
WHERE team_id = ? AND category = ? AND position = ?;

-- name: SetSlot :exec
INSERT INTO team_slots (team_id, category, position, item_id)
VALUES (?, ?, ?, ?)
ON CONFLICT (team_id, category, position)
DO UPDATE SET item_id = excluded.item_id;

-- name: ClearSlot :exec
DELETE FROM team_slots
WHERE team_id = ? AND category = ? AND position = ?;

-- name: SetWeaponSkill :exec
INSERT INTO team_weapon_skills (weapon_skill_id, position, team_id, item_id, slot_category, slot_position)
VALUES (?, ?, ?, ?, ?, ?)
ON CONFLICT
DO UPDATE SET weapon_skill_id = excluded.weapon_skill_id;

-- name: ClearWeaponSkill :exec
DELETE FROM team_weapon_skills
WHERE team_id = ? AND slot_category = ? AND slot_position = ? AND position = ?;

-- name: ClearSlotWeaponSkills :exec
DELETE FROM team_weapon_skills
WHERE team_id = ? AND slot_category = ? AND slot_position = ?;

-- name: GetTeamWeaponsSkills :many
SELECT tws.slot_position, tws.position, ws.name, ws.img
FROM team_weapon_skills tws
JOIN weapon_skills ws ON tws.weapon_skill_id = ws.id
WHERE tws.team_id = ?
ORDER BY tws.slot_position, tws.position;
