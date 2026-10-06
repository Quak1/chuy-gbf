-- name: GetAllItems :many
SELECT * FROM items;

-- name: CreateItem :exec
INSERT INTO items (id, name, element, type, series, enabled)
VALUES (?, ?, ?, ?, ?, ?)
ON CONFLICT(id) DO NOTHING;

-- name: EnableItem :exec
UPDATE items
SET enabled = 1
WHERE id = ?;

-- name: DisableItem :exec
UPDATE items
SET enabled = 0
WHERE id = ?;
