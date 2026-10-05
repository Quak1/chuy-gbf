-- name: GetAllItems :many
SELECT * FROM items;

-- name: CreateItem :exec
INSERT INTO items (id, name, element, type, series, enabled)
VALUES (?, ?, ?, ?, ?, ?)
ON CONFLICT(id) DO NOTHING;
