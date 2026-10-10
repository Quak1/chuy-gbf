-- +goose Up
CREATE TABLE IF NOT EXISTS db_data_updates (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  is_ok BOOLEAN NOT NULL,
  comments TEXT NOT NULL,
  created_at DATETIME DEFAULT (datetime('now')) NOT NULL
) ;

-- +goose Down
DROP TABLE db_data_updates;
