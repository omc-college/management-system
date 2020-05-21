BEGIN;

ALTER TABLE users
    ADD IF NOT EXISTS mobile_phone VARCHAR(255),
    ADD IF NOT EXISTS created_at timestamptz,
    ADD IF NOT EXISTS modified_at timestamptz;

ALTER TABLE users
    ADD CONSTRAINT mobile_phone
        UNIQUE (mobile_phone);

COMMIT;