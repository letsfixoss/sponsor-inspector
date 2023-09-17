package main

import (
	"context"
	"log"

	"github.com/letsfixoss/gh-sponsor-grabber/db"
	"github.com/letsfixoss/gh-sponsor-grabber/repositories"
)

func main() {
	ctx := context.Background()

	repos := repositories.GetRepos()
	log.Printf("Repos: %d", len(repos))

	conn := db.GetConnection()

	for _, repo := range repos {
		log.Printf("Repo: %s", repo.FullName)
		dbRepo := db.Repository{
			Name:      repo.FullName,
			Owner:     repo.Owner.Login,
			AvatarURL: repo.Owner.AvatarURL,
		}
		if err := conn.UpsertRepository(ctx, &dbRepo); err != nil {
			log.Printf("Failed to upsert repository %s: %s", repo.FullName, err)
		}
	}
}
