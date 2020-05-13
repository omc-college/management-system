
BEGIN;

CREATE TABLE IF NOT EXISTS password (
                                     id SERIAL PRIMARY KEY,
                                     current_password VARCHAR(255) NOT NULL,
                                     new_password VARCHAR(255) NOT NULL,
                                     token VARCHAR(255) NOT NULL,
                                     new_token VARCHAR(255) NOT NULL,
                                     verified BOOLEAN DEFAULT false NOT NULL
);
COMMIT;