package db

import (
	"context"
	"fmt"
)

type Repository struct {
	ID        uint64
	Name      string
	Owner     string
	AvatarURL string
}

// UpsertRepository inserts or updates a repository
func (c *Connection) UpsertRepository(ctx context.Context, r *Repository) error {
	const query = `
		INSERT INTO repositories (name, owner, avatar_url) 
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
