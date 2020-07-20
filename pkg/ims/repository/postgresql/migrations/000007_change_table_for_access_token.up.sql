BEGIN;


ALTER TABLE user_acsess_tokens RENAME TO user_access_tokens;

COMMIT;