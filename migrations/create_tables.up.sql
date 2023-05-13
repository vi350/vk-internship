CREATE TABLE IF NOT EXISTS users
(
    id         BIGSERIAL    NOT NULL PRIMARY KEY,
    first_name VARCHAR(100) NOT NULL,
    username   VARCHAR(100),
    start_date BIGINT       NOT NULL,
    language   VARCHAR(2)   NOT NULL,
    state      INT          NOT NULL,
    refer      varchar(100)
);

CREATE TABLE IF NOT EXISTS games
(
    id           SERIAL NOT NULL PRIMARY KEY,
    owner_id     BIGSERIAL REFERENCES users (id),
    opponent_id  BIGSERIAL REFERENCES users (id),
    white_pieces JSONB  NOT NULL,
    black_pieces JSONB  NOT NULL,
    notation     TEXT,
    state        INT    NOT NULL
);