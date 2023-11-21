BEGIN;
ALTER TABLE users
    DROP COLUMN password;
ALTER TABLE users
    ADD COLUMN password bytea;
ALTER TABLE users
    ADD COLUMN salt bytea;
COMMIT;