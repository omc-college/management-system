BEGIN;

CREATE TABLE IF NOT EXISTS users (
<<<<<<< HEAD
    id  SERIAL PRIMARY KEY,
    first_name VARCHAR(255) NOT NULL,
    last_name VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL,
=======
    id SERIAL PRIMARY KEY,
    first_name VARCHAR(255) NOT NULL,
    last_name VARCHAR(255) NOT NULL,
	email VARCHAR(255) NOT NULL,
>>>>>>> d8b4b4c0e6f7106fb7300ca14f37fe09382ee674
	verified BOOLEAN DEFAULT false NOT NULL
);

CREATE TABLE IF NOT EXISTS email_verification_tokens (
    id SERIAL PRIMARY KEY REFERENCES users(id),
    verification_token VARCHAR(255) NOT NULL,
	generated_at TIMESTAMP WITHOUT TIME ZONE NOT NULL
);

CREATE TABLE IF NOT EXISTS credentials (
    id SERIAL PRIMARY KEY REFERENCES users(id),
    password_hash TEXT NOT NULL,
	salt VARCHAR(255) NOT NULL,
	updated_at TIMESTAMP WITHOUT TIME ZONE NOT NULL
);

COMMIT;