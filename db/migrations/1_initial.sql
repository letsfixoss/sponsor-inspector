-- +migrate Up
CREATE TABLE repositories (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT NOT NULL UNIQUE,
    owner TEXT NOT NULL,
    avatar_url TEXT NOT NULL
);

-- +migrate Down
DROP TABLE repositories;
