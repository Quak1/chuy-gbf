-- +goose Up
CREATE TABLE IF NOT EXISTS user_items (
  user_id INTEGER NOT NULL,
  item_id TEXT NOT NULL,
  value TEXT NOT NULL,
  PRIMARY KEY(user_id, item_id),
  FOREIGN KEY (user_id) REFERENCES users(id),
  FOREIGN KEY (item_id) REFERENCES items(id)
);

-- +goose Down
DROP TABLE user_items;
