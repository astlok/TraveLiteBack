DROP SCHEMA IF EXISTS hl CASCADE;

CREATE SCHEMA hl;

CREATE TABLE hl.photo
(
    id   BIGSERIAL PRIMARY KEY NOT NULL,
    link TEXT                  NOT NULL
);

CREATE TABLE hl.chat
(
    id   BIGSERIAL PRIMARY KEY NOT NULL,
    name TEXT                  NOT NULL
);

CREATE TABLE hl.users
(
    id       BIGSERIAL PRIMARY KEY NOT NULL,
    phone    TEXT UNIQUE           NOT NULL,
    nickname TEXT UNIQUE           NOT NULL,
    password TEXT                  NOT NULL,
    photo_id BIGINT,
    FOREIGN KEY (photo_id)
        REFERENCES hl.photo (id)
);

CREATE INDEX hl_phone_idx ON hl.users(phone);

CREATE TABLE hl.sessions
(
    session_id TEXT   NOT NULL,
    user_id    BIGINT NOT NULL,
    FOREIGN KEY (user_id)
        REFERENCES hl.users (id)
);

CREATE TABLE hl.messages
(
    id             BIGSERIAL PRIMARY KEY NOT NULL,
    text           TEXT                  NOT NULL,
    user_sender_id BIGINT                NOT NULL,
    send_time      TIMESTAMP             NOT NULL,
    chat_id        BIGINT                NOT NULL,
    photo_id       BIGINT,
    FOREIGN KEY (user_sender_id)
        REFERENCES hl.users (id),
    FOREIGN KEY (chat_id)
        REFERENCES hl.chat (id),
    FOREIGN KEY (photo_id)
        REFERENCES hl.photo (id)
);

CREATE TABLE hl.chat_users
(
    user_id BIGINT NOT NULL,
    chat_id BIGINT NOT NULL,
    FOREIGN KEY (user_id)
        REFERENCES hl.users (id),
    FOREIGN KEY (chat_id)
        REFERENCES hl.chat (id)
);
