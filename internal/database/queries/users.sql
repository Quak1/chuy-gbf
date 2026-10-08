-- name: GetAllUsers :many
SELECT id, username, comment FROM users
WHERE role != "admin";

-- name: CreateUser :one
INSERT INTO users (username, comment, role)
VALUES (?, "", "")
ON CONFLICT(username) DO UPDATE SET username = username
RETURNING id, username, role;

-- name: GetUser :one
SELECT * FROM users
WHERE id = ?;

-- name: GetUserByUsername :one
SELECT * FROM users
WHERE username = ?;

-- name: MakeUserAdmin :exec
UPDATE users
SET role = "admin"
WHERE id = ?;

-- name: RemoveUserRole :exec
UPDATE users
SET role = ""
WHERE id = ?;

-- name: GetUserItems :many
SELECT COALESCE(ui.value, '') AS value, i.id, i.name, i.element, i.type, i.series
FROM items i
LEFT JOIN user_items ui ON i.id = ui.item_id AND ui.user_id = ?
WHERE i.enabled = 1;

-- name: SetUserItem :exec
INSERT INTO user_items (user_id, item_id, value)
VALUES (?, ?, ?)
ON CONFLICT (user_id, item_id) DO UPDATE SET value=excluded.value;

-- name: SetUserComment :exec
UPDATE users
SET comment = ?
WHERE id = ?;
