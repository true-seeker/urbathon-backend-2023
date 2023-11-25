BEGIN;
ALTER TABLE news
    ADD COLUMN address varchar(1028);
ALTER TABLE news
    ADD COLUMN latitude float;
ALTER TABLE news
    ADD COLUMN longitude float;
COMMIT;