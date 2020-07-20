BEGIN;

ALTER TABLE user_access_tokens
RENAME COLUMN acsess_token TO access_token;


COMMIT;