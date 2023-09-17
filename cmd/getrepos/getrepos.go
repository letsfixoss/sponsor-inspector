package main

import (
	"context"
	"log"

	"github.com/letsfixoss/gh-sponsor-grabber/db"
	"github.com/letsfixoss/gh-sponsor-grabber/internal/repositories"
)

func main() {
	ctx := context.Background()

	repos := repositories.GetRepos()
	log.Printf("Repos: %d", len(repos))

	conn := db.GetConnection()

	for _, repo := range repos {
		log.Printf("Repo: %s", repo.FullName)
		dbRepo := db.Repository{
			Name:       repo.FullName,
			AvatarURL:  repo.Owner.AvatarURL,
			URL:        repo.URL,
			GithubID:   repo.ID,
			Disabled:   repo.Disabled,
			Archived:   repo.Archived,
			OpenIssues: repo.OpenIssuesCount,
			Stars:      repo.StargazersCount,
			Forks:      repo.ForksCount,
			Watchers:   repo.WatchersCount,
			Language:   repo.Language,
			CreatedAt:  repo.CreatedAsTime(),
		}
		repoOwner := db.RepoOwner{
			Name:      repo.Owner.Login,
			GithubID:  repo.Owner.ID,
			AvatarURL: repo.Owner.AvatarURL,
			URL:       repo.Owner.URL,
		}

		ownerID, err := conn.GetOrCreateOwner(ctx, &repoOwner)
		if err != nil {
			log.Printf("Failed to get or create owner %s: %s", repo.Owner.Login, err)
			panic(err)
		}

		dbRepo.Owner = ownerID

		if err := conn.UpsertRepository(ctx, &dbRepo); err != nil {
			log.Printf("Failed to upsert repository %s: %s", repo.FullName, err)
			panic(err)
		}
	}
}
