CREATE TABLE appeals
(
    id             serial primary key,
    user_id        int references users (id)        NOT NULL,
    appeal_type_id int references appeal_types (id) NOT NULL,
    title          varchar(255)                     NOT NULL,
    description    text                             NOT NULL,
    address        text                             NOT NULL,
    x              float                            NOT NULL,
    y              float                            NOT NULL
)