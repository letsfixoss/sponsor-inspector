package db

import (
	"context"
	"fmt"
)

type Repository struct {
	ID   uint64
	Name string
	URL string
}

// UpsertRepository inserts or updates a repository
func (c *Connection) UpsertRepository(ctx context.Context, r *Repository) error {
	const query = `
		INSERT INTO repositories (name) 
			VALUES (?) ON CONFLICT(name) 
			DO UPDATE SET name = name 
			RETURNING id`

	var id uint64

	if err := c.db.QueryRow(query, r.Name).Scan(&id); err != nil {
		return fmt.Errorf("failed to upsert repository %s: %s", r.Name, err)
	}

	r.ID = id

	return nil
}
