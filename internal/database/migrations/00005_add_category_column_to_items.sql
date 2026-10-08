-- +goose Up
ALTER TABLE items
ADD COLUMN category TEXT NOT NULL
DEFAULT 'weapon'
CHECK (category in ('weapon', 'summon', 'character'));

-- +goose Down
ALTER TABLE items DROP COLUMN category;
