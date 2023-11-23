CREATE TABLE appeal_photos
(
    id        serial primary key,
    appeal_id int references appeals (id),
    url       varchar(256)
);