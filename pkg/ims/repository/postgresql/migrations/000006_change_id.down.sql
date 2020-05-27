BEGIN;

SELECT MAX(id)+1 FROM credentials;
CREATE SEQUENCE cred_id_seq MINVALUE 1;
ALTER TABLE credentials ALTER id SET DEFAULT nextval('cred_id_seq');
ALTER SEQUENCE cred_id_seq OWNED BY credentials.id;

SELECT MAX(id)+1 FROM email_verification_tokens;
CREATE SEQUENCE token_id_seq MINVALUE 1;
ALTER TABLE email_verification_tokens ALTER id SET DEFAULT nextval('token_id_seq');
ALTER SEQUENCE token_id_seq OWNED BY email_verification_tokens.id;

UPDATE users SET mobile_phone = NULL WHERE mobile_phone = '';
CREATE UNIQUE INDEX mobile_phone ON users (mobile_phone);

COMMIT;