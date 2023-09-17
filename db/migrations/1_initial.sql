-- +migrate Up
CREATE TABLE repositories (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT NOT NULL UNIQUE
);

-- +migrate Down
DROP TABLE repositories;
