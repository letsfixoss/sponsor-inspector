package db

import (
	"context"
	"fmt"
	"time"
)

type RepoOwner struct {
	ID           uint64
	Name         string
	SponsorCount *uint
	GithubID     uint64
	AvatarURL    string
	URL          string
}

type Repository struct {
	ID         uint64
	Name       string
	Owner      uint64
	AvatarURL  string
	Stars      int
	Forks      int
	Watchers   int
	Language   string
	CreatedAt  time.Time
	URL        string
	OpenIssues int
	Archived   bool
	Disabled   bool
	GithubID   uint64
}

// UpsertRepository inserts or updates a repository
func (c *Connection) UpsertRepository(ctx context.Context, r *Repository) error {
	const query = `
		INSERT INTO repositories (
			name, owner_id, avatar_url, stars, forks, watchers, language, created_at, url, 
			open_issues, archived, disabled, github_id
		) 
			VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
			ON CONFLICT(name)
			DO UPDATE SET name = name 
			RETURNING id`

	var id uint64

	err := c.db.QueryRow(
		query, r.Name, r.Owner, r.AvatarURL, r.Stars, r.Forks, r.Watchers, r.Language, r.CreatedAt,
		r.URL, r.OpenIssues, r.Archived, r.Disabled, r.GithubID,
	).Scan(&id)
	if err != nil {
		return fmt.Errorf("failed to upsert repository %s: %s", r.Name, err)
	}

	r.ID = id

	return nil
}

func (c *Connection) GetOrCreateOwner(ctx context.Context, ro *RepoOwner) (uint64, error) {
	const query = `INSERT INTO repo_owners (name, github_id, avatar_url, url) VALUES (?, ?, ?, ?)
		ON CONFLICT(name) DO UPDATE SET name = name 
		RETURNING id`

	var id uint64

	if err := c.db.QueryRow(query, ro.Name, ro.GithubID, ro.AvatarURL, ro.URL).Scan(&id); err != nil {
		return 0, fmt.Errorf("failed to get or create owner %s: %s", ro.Name, err)
	}

	ro.ID = id

	return id, nil
}

func (c *Connection) GetRepositories(ctx context.Context) ([]*Repository, error) {
	const query = `SELECT id, name, owner_id, avatar_url FROM repositories`

	rows, err := c.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("failed to get repositories: %s", err)
	}

	defer rows.Close()

	var repos []*Repository

	for rows.Next() {
		var r Repository

		if err := rows.Scan(&r.ID, &r.Name, &r.Owner, &r.AvatarURL); err != nil {
			return nil, fmt.Errorf("failed to scan repository: %s", err)
		}

		repos = append(repos, &r)
	}

	return repos, nil
}

// GetRepoOwners returns all repo owners
func (c *Connection) GetRepoOwners(ctx context.Context) ([]*RepoOwner, error) {
	const query = `SELECT id, name, sponsor_count FROM repo_owners`

	rows, err := c.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("failed to get repo owners: %s", err)
	}

	defer rows.Close()

	var owners []*RepoOwner

	for rows.Next() {
		var o RepoOwner

		if err := rows.Scan(&o.ID, &o.Name, &o.SponsorCount); err != nil {
			return nil, fmt.Errorf("failed to scan repo owner: %s", err)
		}

		owners = append(owners, &o)
	}

	return owners, nil
}

func (c *Connection) UpdateSponsorCount(ctx context.Context, id uint64, count *uint) error {
	const query = `UPDATE repo_owners SET sponsor_count = ? WHERE id = ?`

	if _, err := c.db.Exec(query, count, id); err != nil {
		return fmt.Errorf("failed to update sponsor count for repo owner %d: %s", id, err)
	}

	return nil
}
