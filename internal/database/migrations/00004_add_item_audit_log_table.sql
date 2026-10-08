-- +goose Up
CREATE TABLE IF NOT EXISTS item_audit_log (
  id INTEGER PRIMARY KEY,
  user_id INTEGER NOT NULL,
  item_id INTEGER NOT NULL,
  old_value TEXT NOT NULL,
  new_value TEXT NOT NULL,
  created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
  FOREIGN KEY (user_id) REFERENCES users(id),
  FOREIGN KEY (item_id) REFERENCES items(id)
);

-- +goose Down
DROP TABLE item_audit_log;
