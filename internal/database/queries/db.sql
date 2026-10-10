-- name: AddDBDataUpdate :exec
INSERT INTO db_data_updates (is_ok, comments)
VALUES (?, ?);

-- name: GetAllDBDataUpdates :many
SELECT * FROM db_data_updates;

-- name: GetLastDBDataUpdate :one
SELECT * FROM db_data_updates
WHERE is_ok = 1
ORDER BY id DESC
LIMIT 1;
