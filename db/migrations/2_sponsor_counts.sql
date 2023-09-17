-- +migrate Up
ALTER TABLE repo_owners ADD COLUMN sponsor_count INTEGER NULL DEFAULT NULL;
-- +migrate Down 
ALTER TABLE repo_owners DROP COLUMN sponsor_count;
