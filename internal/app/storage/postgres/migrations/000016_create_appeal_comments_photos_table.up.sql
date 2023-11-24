CREATE TABLE appeal_comment_photos
(
    id                serial primary key,
    appeal_comment_id int references appeal_comments (id),
    url               varchar(256)
);