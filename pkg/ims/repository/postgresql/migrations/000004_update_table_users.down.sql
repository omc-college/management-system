BEGIN;

ALTER TABLE users
    DROP COLUMN IF EXISTS mobile_phone CASCADE,
    DROP COLUMN IF EXISTS created_at CASCADE,
    DROP COLUMN IF EXISTS modified_at CASCADE;

COMMIT;