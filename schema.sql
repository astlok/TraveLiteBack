ALTER
    USER postgres WITH ENCRYPTED PASSWORD 'admin';
DROP SCHEMA IF EXISTS travelite CASCADE;
CREATE
    EXTENSION IF NOT EXISTS citext;
CREATE SCHEMA travelite;

CREATE TABLE travelite.users
(
    id       SERIAL PRIMARY KEY NOT NULL,
    email    CITEXT UNIQUE      NOT NULL,
    nickname CITEXT UNIQUE      NOT NULL,
    password TEXT               NOT NULL,
    img      TEXT DEFAULT ''
);

CREATE TABLE travelite.sessions
(
    user_id    INT  NOT NULL,
    auth_token TEXT NOT NULL,
    FOREIGN KEY (user_id)
        REFERENCES travelite.users (id)
);

CREATE TABLE travelite.region
(
    id   SERIAL PRIMARY KEY NOT NULL,
    name CITEXT UNIQUE      NOT NULL
);

CREATE TABLE travelite.trek
(
    id          SERIAL PRIMARY KEY NOT NULL,
    name        CITEXT             NOT NULL,
    difficult   INT                NOT NULL,
    days        INT                NOT NULL,
    description CITEXT             NOT NULL,
    file        TEXT               NOT NULL,
    region_id   INT                NOT NULL,
    user_id     INT                NOT NULL,
    FOREIGN KEY (user_id)
        REFERENCES travelite.users (id),
    FOREIGN KEY (region_id)
        REFERENCES travelite.region (id)
);

CREATE TABLE travelite.trek_rating
(
    user_id INT NOT NULL,
    trek_id INT NOT NULL,
    rating  FLOAT8 NOT NULL,
    FOREIGN KEY (user_id)
        REFERENCES travelite.users (id),
    FOREIGN KEY (trek_id)
        REFERENCES travelite.trek (id)
);

CREATE TABLE travelite.things
(
    id   SERIAL PRIMARY KEY NOT NULL,
    name CITEXT             NOT NULL
);

CREATE TABLE travelite.trek_things
(
    trek_id  INT NOT NULL,
    thing_id INT NOT NULL,
    FOREIGN KEY (trek_id)
        REFERENCES travelite.trek (id) ON DELETE CASCADE,
    FOREIGN KEY (thing_id)
        REFERENCES travelite.things (id) ON DELETE CASCADE
);

CREATE TABLE travelite.comment
(
    id          SERIAL PRIMARY KEY NOT NULL,
    trek_id     INT                NOT NULL,
    user_id     INT                NOT NULL,
    description CITEXT             NOT NULL,
    FOREIGN KEY (user_id)
        REFERENCES travelite.users (id)
);

CREATE TABLE travelite.comment_photo
(
    id         SERIAL PRIMARY KEY NOT NULL,
    comment_id INT                NOT NULL,
    trek_id    INT                NOT NULL,
    user_id    INT                NOT NULL,
    photo_url  TEXT               NOT NULL,
    FOREIGN KEY (trek_id)
        REFERENCES travelite.trek (id),
    FOREIGN KEY (comment_id)
        REFERENCES travelite.comment (id),
    FOREIGN KEY (user_id)
        REFERENCES travelite.users (id)
);

INSERT INTO travelite.region (name)
VALUES ('Москва и МО'),
       ('Белгородская область'),
       ('Брянская область'),
       ('Владимирская область'),
       ('Воронежская область'),
       ('Ивановская область'),
       ('Калужская область'),
       ('Костромская область'),
       ('Курская область'),
       ('Липецкая область'),
       ('Орловская область'),
       ('Рязанская область'),
       ('Смоленская область'),
       ('Тамбовская область'),
       ('Тверская область'),
       ('Тульская область'),
       ('Ярославская область'),
       ('Республика Карелия'),
       ('Республика Коми'),
       ('Архангельская область'),
       ('Ненецкий автономный округ'),
       ('Вологодская область'),
       ('Калининградская область'),
       ('Ленинградская область'),
       ('Мурманская область'),
       ('Новгородская область'),
       ('Псковская область'),
       ('Республика Адыгея'),
       ('Республика Дагестан'),
       ('Республика Ингушетия'),
       ('Кабардино-Балкарская Республика'),
       ('Республика Калмыкия'),
       ('Карачаево-Черкесская Республика'),
       ('Республика Северная Осетия - Алания'),
       ('Чеченская Республика'),
       ('Краснодарский край'),
       ('Ставропольский край'),
       ('Астраханская область'),
       ('Волгоградская область'),
       ('Ростовская область'),
       ('Республика Башкортостан'),
       ('Республика Марий Эл'),
       ('Республика Мордовия'),
       ('Республика Татарстан'),
       ('Удмуртская Республика'),
       ('Чувашская Республика'),
       ('Кировская область'),
       ('Нижегородская область'),
       ('Оренбургская область'),
       ('Пензенская область'),
       ('Пермская область'),
       ('Коми-Пермяцкий автономный округ'),
       ('Самарская область'),
       ('Саратовская область'),
       ('Ульяновская область'),
       ('Курганская область'),
       ('Свердловская область'),
       ('Тюменская область'),
       (' Ханты-Мансийский автономный округ - Югра'),
       ('Ямало-Ненецкий автономный округ'),
       ('Челябинская область'),
       ('Республика Алтай'),
       ('Республика Бурятия'),
       ('Республика Тыва'),
       ('Республика Хакасия'),
       ('Алтайский край'),
       ('Красноярский край'),
       ('Таймырский автономный округ'),
       ('Эвенкийский автономный округ'),
       ('Иркутская область'),
       ('Кемеровская область'),
       ('Новосибирская область'),
       ('Омская область'),
       ('Томская область'),
       ('Читинская область'),
       ('Агинский Бурятский автономный округ'),
       ('Республика Саха (Якутия)'),
       ('Приморский край'),
       ('Хабаровский край'),
       ('Амурская область'),
       ('Камчатская область'),
       ('Корякский автономный округ'),
       ('Магаданская область'),
       ('Сахалинская область'),
       ('Еврейская автономная область'),
       ('Чукотский автономный округ')
