BEGIN;

CREATE TABLE IF NOT EXISTS resources (
    id  SERIAL PRIMARY KEY,
    resource_name VARCHAR(255) NOT NULL,
    resource_description VARCHAR(1024) NOT NULL,
    modified_at timestamptz NOT NULL
);

CREATE TABLE IF NOT EXISTS rooms (
    id  SERIAL PRIMARY KEY,
    room VARCHAR(255) NOT NULL,
    modified_at timestamptz NOT NULL
);

CREATE TABLE IF NOT EXISTS subjects (
    id  SERIAL PRIMARY KEY,
    subject_name VARCHAR(255) NOT NULL,
    modified_at timestamptz NOT NULL
);

COMMIT;