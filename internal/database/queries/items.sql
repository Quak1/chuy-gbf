-- name: GetAllItems :many
SELECT * FROM items;

-- name: CreateItem :exec
INSERT OR REPLACE INTO items (id, name, element, type, series, enabled, category)
VALUES (?, ?, ?, ?, ?, ?, ?);

-- name: EnableItem :exec
UPDATE items
SET enabled = 1
WHERE id = ?;

-- name: DisableItem :exec
UPDATE items
SET enabled = 0
WHERE id = ?;

-- name: GetEnabledItems :many
SELECT * FROM items
WHERE enabled = 1;

-- name: GetUserItemValues :many
SELECT * FROM user_items;
