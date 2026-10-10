-- name: GetAllUsers :many
SELECT id, username, comment FROM users
WHERE role != "admin";

-- name: CreateUser :one
INSERT INTO users (username, comment, role)
VALUES (?, "", ?)
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
SELECT
  i.*,
  COALESCE(iv.id, 0) AS item_value_id,
  COALESCE(iv.value, '') AS value,
  COALESCE(iv.color, '') AS color
FROM items i
LEFT JOIN user_items ui ON ui.item_id = i.id AND ui.user_id = ?
LEFT JOIN item_values iv ON iv.id = ui.item_value_id
WHERE i.enabled = 1;

-- name: SetUserItem :exec
INSERT INTO user_items (user_id, item_id, item_value_id)
VALUES (?, ?, ?)
ON CONFLICT DO UPDATE SET item_value_id = excluded.item_value_id;

-- name: SetUserComment :exec
UPDATE users
SET comment = ?
WHERE id = ?;

-- name: DeleteUser :exec
DELETE FROM users
WHERE id = ?;
