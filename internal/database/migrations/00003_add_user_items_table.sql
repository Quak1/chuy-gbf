-- +goose Up
CREATE TABLE IF NOT EXISTS user_items (
  user_id INTEGER NOT NULL,
  item_id TEXT NOT NULL,
  item_value_id INTEGER NOT NULL,
  PRIMARY KEY(user_id, item_id),
  FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
  FOREIGN KEY (item_value_id, item_id) REFERENCES item_values(id, item_id) ON DELETE CASCADE
);

-- +goose Down
DROP TABLE user_items;
