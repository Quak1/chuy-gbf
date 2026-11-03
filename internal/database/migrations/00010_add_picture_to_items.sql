-- +goose Up
ALTER TABLE items
ADD COLUMN image_name TEXT NOT NULL
DEFAULT '';

-- +goose Down
ALTER TABLE items DROP COLUMN image_name;
