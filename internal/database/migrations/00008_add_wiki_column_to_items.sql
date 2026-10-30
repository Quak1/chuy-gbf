-- +goose Up
ALTER TABLE items
ADD COLUMN wiki_path TEXT NOT NULL
DEFAULT '';

-- +goose Down
ALTER TABLE items DROP COLUMN wiki_path;
