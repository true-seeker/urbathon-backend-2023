CREATE TABLE appeal_types
(
    id              serial primary key,
    title           varchar(256),
    appeal_theme_id int references appeal_themes (id)
)