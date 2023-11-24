BEGIN;
CREATE TABLE news_categories
(
    id    serial primary key,
    title varchar(64)
);

INSERT INTO news_categories(id, title)
VALUES (1, 'Вода'),
       (2, 'Свет'),
       (3, 'Газ'),
       (4, 'Отопление'),
       (5, 'Улица'),
       (6, 'Отходы'),
       (7, 'Дороги'),
       (8, 'Зеленые зоны'),
       (9, 'Безопасность'),
       (10, 'Транспорт'),
       (11, 'Экология')
;

ALTER TABLE news
    ADD COLUMN category_id int references news_categories (id) DEFAULT 1;

COMMIT;