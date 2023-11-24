BEGIN;
ALTER TABLE news
    ADD COLUMN user_id int references users (id);
ALTER TABLE news
    ADD COLUMN organization_id int references organizations (id);
COMMIT;