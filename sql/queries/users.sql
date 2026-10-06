-- name: GetAllUsers :many
SELECT * FROM users;

-- name: CreateUser :exec
INSERT INTO users (username, comment, role)
VALUES (?, ?, ?);

-- name: GetUser :one
SELECT * FROM users
WHERE id = ?;

-- name: MakeUserAdmin :exec
UPDATE users
SET role = "admin"
WHERE id = ?;

-- name: RemoveUserRole :exec
UPDATE users
SET role = ""
WHERE id = ?;

-- name: GetUserItems :many
SELECT ui.item_id, ui.value, i.name, i.element, i.type, i.series
FROM users u
JOIN user_items ui ON u.id = ui.user_id
JOIN items i ON i.id = ui.item_id
WHERE u.id = ? AND i.enabled = 1;

-- name: SetUserItem :exec
INSERT INTO user_items (user_id, item_id, value)
VALUES (?, ?, ?)
ON CONFLICT (user_id, item_id) DO UPDATE SET value=excluded.value;
