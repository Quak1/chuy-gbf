-- +goose Up
CREATE TABLE IF NOT EXISTS item_values (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  item_id TEXT NOT NULL,
  value TEXT NOT NULL,
  color TEXT NOT NULL CHECK (
        LENGTH(color) = 7 AND
        color GLOB '#[0-9a-fA-F][0-9a-fA-F][0-9a-fA-F][0-9a-fA-F][0-9a-fA-F][0-9a-fA-F]'
    ),
  UNIQUE(item_id, value),
  UNIQUE(id, item_id),
  FOREIGN KEY (item_id) REFERENCES items(id)
);

-- +goose Down
DROP TABLE item_values;
