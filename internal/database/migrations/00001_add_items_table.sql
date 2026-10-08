-- +goose Up
CREATE TABLE IF NOT EXISTS items (
  id TEXT PRIMARY KEY NOT NULL,
  name TEXT NOT NULL,
  element TEXT NOT NULL CHECK (element in ('fire', 'water', 'earth', 'wind', 'light', 'dark', 'any')),
  type TEXT NOT NULL,
  series TEXT NOT NULL,
  enabled BOOLEAN NOT NULL DEFAULT 0
) WITHOUT ROWID;

-- +goose Down
DROP TABLE items;
