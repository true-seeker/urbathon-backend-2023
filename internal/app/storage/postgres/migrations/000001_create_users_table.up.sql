CREATE TABLE users
(
    id       serial primary key,
    name     varchar(255),
    email    varchar(128),
    password varchar(128)
);