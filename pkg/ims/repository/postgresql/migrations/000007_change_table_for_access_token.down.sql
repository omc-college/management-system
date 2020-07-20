BEGIN;

ALTER TABLE user_access_tokens RENAME TO user_acsess_tokens;

COMMIT;