BEGIN;

ALTER TABLE user_access_tokens
CHANGE COLUMN  acsess_token access_token VARCHAR(255) NOT NULL;


COMMIT;