BEGIN;

CREATE TABLE IF NOT EXISTS users (
         id  SERIAL PRIMARY KEY,
         firstName VARCHAR(255) NOT NULL,
         lastName VARCHAR(255) NOT NULL,
         email VARCHAR(255) NOT NULL,
         mobilePhone VARCHAR(255) NOT NULL,
         createdAt timestamptz,
         modifiedAt timestamptz
);

COMMIT; 