CREATE TABLE IF NOT EXISTS sets (
    id INTEGER PRIMARY KEY,
    name TEXT NOT NULL,
    slug TEXT NOT NULL,
    type TEXT NOT NULL,
    collectible_count INTEGER NOT NULL,
    collectible_revealed_count INTEGER NOT NULL,
    non_collectible_count INTEGER NOT NULL,
    non_collectible_revealed_count INTEGER NOT NULL,
    alias_set_ids INTEGER[] NOT NULL DEFAULT '{}'
);

CREATE TABLE IF NOT EXISTS classes (
    id INTEGER PRIMARY KEY,
    slug TEXT NOT NULL,
    name TEXT NOT NULL,
    card_id INTEGER NOT NULL,
    hero_power_card_id INTEGER NOT NULL,
    alternate_hero_card_ids INTEGER[] NOT NULL DEFAULT '{}'
);

CREATE TABLE IF NOT EXISTS rarities (
    id INTEGER PRIMARY KEY,
    slug TEXT NOT NULL,
    crafting_cost INTEGER[] NOT NULL DEFAULT '{}',
    dust_value INTEGER[] NOT NULL DEFAULT '{}',
    name TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS types (
    id INTEGER PRIMARY KEY,
    slug TEXT NOT NULL,
    name TEXT NOT NULL,
    game_modes INTEGER[] NOT NULL DEFAULT '{}'
);

CREATE TABLE IF NOT EXISTS keywords (
    id INTEGER PRIMARY KEY,
    slug TEXT NOT NULL,
    name TEXT NOT NULL,
    ref_text TEXT NOT NULL,
    text TEXT NOT NULL,
    game_modes INTEGER[] NOT NULL DEFAULT '{}'
);

CREATE TABLE IF NOT EXISTS cards (
    id INTEGER PRIMARY KEY,
    collectible INTEGER NOT NULL,
    slug TEXT NOT NULL,
    classid INTEGER NOT NULL,
    multiclassids JSONB NOT NULL DEFAULT '[]'::jsonb,
    spellschoolid INTEGER NOT NULL DEFAULT 0,
    cardtypeid INTEGER NOT NULL,
    cardsetid INTEGER NOT NULL,
    rarityid INTEGER NOT NULL,
    artistname TEXT NOT NULL,
    health INTEGER NOT NULL,
    attack INTEGER NOT NULL,
    manacost INTEGER NOT NULL,
    name TEXT NOT NULL,
    text TEXT NOT NULL,
    image TEXT NOT NULL,
    imagegold TEXT NOT NULL,
    flavortext TEXT NOT NULL,
    cropimage TEXT NOT NULL,
    parentid INTEGER NOT NULL,
    copyofcardids INTEGER[] NOT NULL DEFAULT '{}',
    miniontypeid INTEGER NOT NULL DEFAULT 0,
    childids INTEGER[] NOT NULL DEFAULT '{}',
    durability INTEGER NOT NULL DEFAULT 0,
    multitypeids INTEGER[] NOT NULL DEFAULT '{}',
    armor INTEGER NOT NULL DEFAULT 0,
    iszilliaxfunctionalmodule BOOLEAN NOT NULL DEFAULT FALSE,
    iszilliaxcosmeticmodule BOOLEAN NOT NULL DEFAULT FALSE,
    duals_relevant BOOLEAN NOT NULL DEFAULT FALSE,
    duals_constructed BOOLEAN NOT NULL DEFAULT FALSE
);

CREATE TABLE IF NOT EXISTS card_keywords (
    card_id INTEGER NOT NULL REFERENCES cards(id) ON DELETE CASCADE,
    keyword_id INTEGER NOT NULL,
    PRIMARY KEY (card_id, keyword_id)
);

CREATE TABLE IF NOT EXISTS update_meta (
    id BIGSERIAL PRIMARY KEY,
    updated TIMESTAMPTZ NOT NULL,
    is_changed BOOLEAN NOT NULL DEFAULT FALSE
);

CREATE INDEX IF NOT EXISTS idx_cards_mana_name ON cards (manacost, name);
CREATE INDEX IF NOT EXISTS idx_cards_name ON cards (name);
CREATE EXTENSION IF NOT EXISTS pg_trgm;
CREATE INDEX IF NOT EXISTS idx_cards_name_trgm ON cards USING GIN (name gin_trgm_ops);
CREATE INDEX IF NOT EXISTS idx_cards_classid ON cards (classid);
CREATE INDEX IF NOT EXISTS idx_cards_rarityid ON cards (rarityid);
CREATE INDEX IF NOT EXISTS idx_cards_cardtypeid ON cards (cardtypeid);
CREATE INDEX IF NOT EXISTS idx_cards_cardsetid ON cards (cardsetid);
CREATE INDEX IF NOT EXISTS idx_cards_attack ON cards (attack);
CREATE INDEX IF NOT EXISTS idx_cards_health ON cards (health);
CREATE INDEX IF NOT EXISTS idx_card_keywords_keyword_card ON card_keywords (keyword_id, card_id);
CREATE INDEX IF NOT EXISTS idx_card_keywords_card_keyword ON card_keywords (card_id, keyword_id);
