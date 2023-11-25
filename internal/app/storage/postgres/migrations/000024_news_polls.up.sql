BEGIN;
CREATE TABLE news_polls
(
    id    serial primary key,
    title varchar(256)
);
CREATE TABLE poll_options
(
    id      serial primary key,
    poll_id int references news_polls (id),
    title   varchar(256)
);
CREATE TABLE user_poll_votes
(
    id                 serial primary key,
    user_id            int references users (id),
    selected_option_id int references poll_options (id)
);
ALTER TABLE news
    ADD COLUMN poll_id int references news_polls (id);
COMMIT;