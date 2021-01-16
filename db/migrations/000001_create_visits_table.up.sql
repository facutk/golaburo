BEGIN;

CREATE TABLE IF NOT EXISTS visits (
   id serial PRIMARY KEY,
   hits integer
);

INSERT INTO visits (hits) VALUES (0);

COMMIT;