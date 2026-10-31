-- +goose Up
CREATE TABLE teams (
  id         INTEGER  PRIMARY KEY AUTOINCREMENT,
  name       TEXT     NOT NULL,
  url        TEXT     NOT NULL,
  created_at DATETIME NOT NULL DEFAULT (datetime('now')),
  updated_at DATETIME NOT NULL DEFAULT (datetime('now'))
);

CREATE TABLE team_slots (
  team_id  INTEGER NOT NULL,
  item_id  INTEGER NOT NULL,
  category TEXT    NOT NULL CHECK (category IN ('character', 'summon', 'weapon')),
  position INTEGER NOT NULL,

  PRIMARY KEY (team_id, category, position),

  CHECK (
    (category = 'character' AND position BETWEEN 1 AND 6)
    OR
    (category = 'summon' AND position BETWEEN 1 AND 7)
    OR
    (category = 'weapon' AND position BETWEEN 1 AND 10)
  ),

  FOREIGN KEY (team_id) REFERENCES teams(id) ON DELETE CASCADE,
  FOREIGN KEY (item_id, category) REFERENCES items(id, category) ON DELETE CASCADE
) WITHOUT ROWID;

CREATE INDEX team_slots_item_idx ON team_slots(item_id);


-- +goose Down
DROP INDEX IF EXISTS team_slots_item_idx;
DROP TABLE IF EXISTS team_slots;
DROP TABLE IF EXISTS teams;
