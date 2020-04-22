BEGIN;

CREATE TABLE IF NOT EXISTS endpoints (
    id SERIAL PRIMARY KEY,
    path varchar(255) NOT NULL,
    method varchar(255) NOT NULL
);

CREATE TABLE IF NOT EXISTS features (
    id SERIAL PRIMARY KEY,
    name varchar(255) NOT NULL
);

CREATE TABLE IF NOT EXISTS roles (
    id SERIAL PRIMARY KEY,
    name varchar(255) NOT NULL
);

CREATE TABLE IF NOT EXISTS features_to_endpoints (
    feature_id integer NOT NULL REFERENCES features(id),
    endpoint_id integer NOT NULL REFERENCES endpoints(id)
);

CREATE TABLE IF NOT EXISTS roles_to_features (
    role_id integer NOT NULL REFERENCES roles(id) ON UPDATE CASCADE ON DELETE CASCADE,
    feature_id integer NOT NULL REFERENCES features(id) ON UPDATE CASCADE ON DELETE CASCADE
);

COMMIT;