SET SCHEMA :'schema';
--
-- Previously exported data includes serial PK values, so we bump our
-- sequences manually.
--

SELECT setval(pg_get_serial_sequence('game', 'game_id'),
              COALESCE((SELECT MAX(game_id)+1 FROM game), 1), false);

SELECT setval(pg_get_serial_sequence('publisher', 'publisher_id'),
              COALESCE((SELECT MAX(publisher_id)+1 FROM publisher), 1), false);
