BEGIN;

ALTER TABLE user_access_tokens
DROP COLUMN IF EXISTS access_tokens;

COMMIT;