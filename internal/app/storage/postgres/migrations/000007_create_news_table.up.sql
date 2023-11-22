CREATE TABLE news
(
    id    serial primary key,
    title varchar(128) NOT NULL,
    body  varchar      NOT NULL,
    date  timestamp with time zone DEFAULT now()
)