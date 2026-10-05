-- +goose Up
CREATE TABLE IF NOT EXISTS users (
  id INTEGER PRIMARY KEY autoincrement,
  name TEXT NOT NULL UNIQUE,
  comment TEXT NOT NULL,
  role TEXT NOT NULL
);

-- +goose Down
DROP TABLE users;
