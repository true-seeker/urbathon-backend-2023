CREATE TABLE appeal_themes
(
    id                 serial primary key,
    title              varchar(256),
    appeal_category_id int references appeal_categories (id)
)