CREATE TABLE appeal_comments
(
    id        serial primary key,
    appeal_id int references appeals (id),
    user_id   int references users (id),
    date      timestamp with time zone default now(),
    text      varchar
);