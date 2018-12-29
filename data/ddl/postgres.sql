CREATE SCHEMA IF NOT EXISTS :schema;
SET SCHEMA :'schema';

DROP TABLE IF EXISTS "game" CASCADE;
CREATE TABLE "game" (
  game_id serial,
  game varchar(1024) NOT NULL,
  box_art varchar(256),
  na_rel_dt DATE,
  eu_rel_dt DATE,
  loose_price INTEGER,
  licensed BOOLEAN NOT NULL,
  wikipedia_url varchar(1024),
  CONSTRAINT game_pk PRIMARY KEY (game_id)
)
WITH ( OIDS = FALSE );

CREATE INDEX game_lower_game_nm_idx ON "game" (lower(game));


DROP TABLE IF EXISTS "publisher" CASCADE;
CREATE TABLE "publisher" (
  publisher_id SERIAL,
  publisher VARCHAR(1024) UNIQUE NOT NULL,
  wikipedia_url VARCHAR(1024),
  CONSTRAINT publisher_pk PRIMARY KEY (publisher_id)
)
WITH ( OIDS = FALSE );

CREATE INDEX publisher_lower_publisher_nm_idx ON publisher (lower(publisher));


DROP TABLE IF EXISTS "game_x_publisher" CASCADE;
CREATE TABLE "game_x_publisher" (
  game_id INTEGER NOT NULL,
  publisher_id INTEGER NOT NULL,
  CONSTRAINT game_x_publisher_pk PRIMARY KEY (game_id, publisher_id)
)
WITH ( OIDS = FALSE );
