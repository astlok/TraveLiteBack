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
    rating      INT,
    region_id   INT                NOT NULL,
    user_id     INT                NOT NULL,
    FOREIGN KEY (user_id)
        REFERENCES travelite.users (id),
    FOREIGN KEY (region_id)
        REFERENCES travelite.region (id)
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
        REFERENCES travelite.trek (id),
    FOREIGN KEY (thing_id)
        REFERENCES travelite.things (id)
);

CREATE TABLE travelite.comment
(
    id          SERIAL PRIMARY KEY NOT NULL,
    trek_id     INT                NOT NULL,
    user_id     INT                NOT NULL,
    description CITEXT             NOT NULL,
    FOREIGN KEY (user_id)
        REFERENCES travelite.users (id),
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