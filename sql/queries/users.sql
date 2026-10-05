-- name: GetAllUsers :many
SELECT * FROM users;

-- name: CreateUser :exec
INSERT INTO users (name, comment, role)
VALUES (?, ?, ?);

-- name: MakeUserAdmin :exec
UPDATE users
SET role = "admin"
WHERE id = ?;

-- name: RemoveUserRole :exec
UPDATE users
SET role = ""
WHERE id = ?;
