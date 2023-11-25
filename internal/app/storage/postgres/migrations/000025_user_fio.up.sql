BEGIN;
ALTER TABLE users
    DROP COLUMN name;
ALTER TABLE users
    ADD COLUMN first_name varchar(128);
ALTER TABLE users
    ADD COLUMN last_name varchar(128);
ALTER TABLE users
    ADD COLUMN patronymic varchar(128);
COMMIT;