-- +migrate Up
CREATE TABLE repo_owners (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT NOT NULL UNIQUE
);

CREATE TABLE repositories (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT NOT NULL UNIQUE,
    owner_id INTEGER NOT NULL references repo_owners(id),
    avatar_url TEXT NOT NULL
);

-- +migrate Down
DROP TABLE repositories;
DROP TABLE repo_owners;
