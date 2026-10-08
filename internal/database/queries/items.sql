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

-- name: GetAllSelectedItemValues :many
SELECT ui.user_id, iv.id AS value_id, iv.item_id, iv.value, iv.color FROM user_items ui
JOIN item_values iv ON ui.item_value_id = iv.id;

-- name: CreateItemValue :one
INSERT INTO item_values  (item_id, value, color)
VALUES (?, ?, ?)
RETURNING *;

-- name: DeleteItemValue :exec
DELETE FROM item_values WHERE id = ? AND item_id = ?;

-- name: GetItemValues :many
SELECT * FROM item_values
WHERE item_id = ?;
