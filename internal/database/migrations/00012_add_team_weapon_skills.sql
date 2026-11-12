-- +goose Up
CREATE UNIQUE INDEX weapon_skills_id_weapon_id_uq ON weapon_skills (id, weapon_id);
CREATE UNIQUE INDEX team_slots_teamid_itemid_category_position ON team_slots (team_id, item_id, category, position);

CREATE TABLE IF NOT EXISTS team_weapon_skills (
  weapon_skill_id INTEGER NOT NULL,
  position INTEGER    NOT NULL CHECK (position BETWEEN 1 AND 3),

  team_id INTEGER NOT NULL,
  item_id TEXT NOT NULL,
  slot_category TEXT    NOT NULL CHECK (slot_category = 'weapon'),
  slot_position INTEGER NOT NULL,

  PRIMARY KEY (team_id, slot_category, slot_position, position),
  -- UNIQUE (team_id, slot_category, slot_position, weapon_skill_id)

  FOREIGN KEY (team_id, item_id, slot_category, slot_position)
    REFERENCES team_slots(team_id, item_id, category, position)
    ON DELETE CASCADE,

  FOREIGN KEY (weapon_skill_id, item_id)
    REFERENCES weapon_skills(id, weapon_id)
);

-- +goose Down
DROP TABLE team_weapon_skills;
DROP INDEX team_slots_teamid_itemid_category_position;
DROP INDEX weapon_skills_id_weapon_id_uq;
