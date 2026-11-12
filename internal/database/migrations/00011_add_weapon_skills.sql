-- +goose Up
CREATE TABLE IF NOT EXISTS weapon_skills (
  id INTEGER PRIMARY KEY,
  weapon_id TEXT NOT NULL,
  name TEXT NOT NULL,
  img TEXT NOT NULL,

  FOREIGN KEY (weapon_id) REFERENCES items(id) ON DELETE CASCADE,
  UNIQUE (weapon_id, name)
);

CREATE INDEX weapon_skills_weapon_idx ON weapon_skills(weapon_id);

-- +goose Down
DROP INDEX IF EXISTS weapon_skills_weapon_idx;
DROP TABLE weapon_skills;
