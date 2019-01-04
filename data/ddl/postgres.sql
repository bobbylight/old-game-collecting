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
  CONSTRAINT game_x_publisher_pk PRIMARY KEY (game_id, publisher_id),
  CONSTRAINT game_x_publisher_game_fk FOREIGN KEY (game_id)
    REFERENCES game (game_id) ON UPDATE NO ACTION ON DELETE NO ACTION,
  CONSTRAINT game_x_publisher_publisher_fk FOREIGN KEY (publisher_id)
    REFERENCES publisher (publisher_id) ON UPDATE NO ACTION ON DELETE NO ACTION
)
WITH ( OIDS = FALSE );


DROP TABLE IF EXISTS "user_x_game" CASCADE;
CREATE TABLE "user_x_game" (
  user_id INTEGER NOT NULL,
  game_id INTEGER NOT NULL,
  CONSTRAINT "user_x_game_pk" PRIMARY KEY (user_id, game_id),
  CONSTRAINT user_x_game_game_fk FOREIGN KEY (game_id)
    REFERENCES game (game_id) ON UPDATE NO ACTION ON DELETE NO ACTION
)
WITH ( OIDS = FALSE );
