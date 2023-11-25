create table tko
(
    id        serial primary key,
    address   text,
    latitude  double precision,
    longitude double precision,
    type      varchar(256)
);