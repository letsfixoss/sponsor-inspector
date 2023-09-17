package db

import (
	"context"
	"fmt"
)

type RepoOwners struct {
	ID   uint64
	Name string
}

type Repository struct {
	ID        uint64
	Name      string
	Owner     uint64
	AvatarURL string
}

// UpsertRepository inserts or updates a repository
func (c *Connection) UpsertRepository(ctx context.Context, r *Repository) error {
	const query = `
		INSERT INTO repositories (name, owner_id, avatar_url) 
			VALUES (?, ?, ?) ON CONFLICT(name) 
			DO UPDATE SET name = name 
			RETURNING id`

	var id uint64

	if err := c.db.QueryRow(query, r.Name, r.Owner, r.AvatarURL).Scan(&id); err != nil {
		return fmt.Errorf("failed to upsert repository %s: %s", r.Name, err)
	}

	r.ID = id

	return nil
}

func (c *Connection) GetOrCreateOwner(ctx context.Context, name string) (uint64, error) {
	const query = `INSERT INTO repo_owners (name) VALUES (?) ON CONFLICT(name) DO UPDATE SET name = name RETURNING id`

	var id uint64

	if err := c.db.QueryRow(query, name).Scan(&id); err != nil {
		return 0, fmt.Errorf("failed to get or create owner %s: %s", name, err)
	}

	return id, nil
}
