BEGIN;
CREATE TABLE organizations
(
    id      serial primary key,
    name    varchar(512),
    inn     varchar(32),
    address varchar(1024),
    phone   varchar(16)
);
CREATE TABLE organization_appeal_category
(
    id                 serial primary key,
    organization_id    int references organizations (id),
    appeal_category_id int references appeal_categories (id)
);
ALTER TABLE users
    ADD COLUMN organization_id int references organizations (id);

COMMIT;