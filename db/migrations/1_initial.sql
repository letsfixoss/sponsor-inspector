-- +migrate Up
CREATE TABLE repo_owners (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT NOT NULL UNIQUE,
    sponsor_count INTEGER NULL DEFAULT NULL,
    github_id INTEGER NOT NULL UNIQUE,
    avatar_url TEXT NOT NULL,
    url TEXT NOT NULL
);

CREATE TABLE repositories (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT NOT NULL UNIQUE,
    owner_id INTEGER NOT NULL references repo_owners(id),
    avatar_url TEXT NOT NULL,
    stars INTEGER NOT NULL,
    forks INTEGER NOT NULL,
    watchers INTEGER NOT NULL,
    language TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL,
    github_id INTEGER NOT NULL UNIQUE,
    url TEXT NOT NULL,
    open_issues INTEGER NOT NULL,
    archived BOOLEAN NOT NULL,
    disabled BOOLEAN NOT NULL
);

-- +migrate Down
DROP TABLE repositories;
DROP TABLE repo_owners;
