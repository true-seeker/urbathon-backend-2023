BEGIN;
CREATE TABLE appeal_status
(
    id     serial primary key,
    status varchar(32)
);

INSERT INTO appeal_status(id, status)
VALUES (1, 'Запланировано'),
       (2, 'В работе'),
       (3, 'Подтверждается'),
       (4, 'Решено'),
       (5, 'Отклонено');

ALTER TABLE appeals
    ADD COLUMN status_id int references appeal_status (id) NOT NULL DEFAULT 1;
COMMIT;