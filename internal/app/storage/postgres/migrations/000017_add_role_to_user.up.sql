BEGIN;
CREATE TABLE user_roles
(
    id    serial primary key,
    title varchar(64)
);

INSERT INTO user_roles (id, title)
VALUES (1, 'User'),
       (2, 'Municipal service'),
       (3, 'Admin');

ALTER TABLE users
    ADD COLUMN role int references user_roles (id) DEFAULT 1;

COMMIT;