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
			Name:      repo.FullName,
			AvatarURL: repo.Owner.AvatarURL,
		}

		ownerID, err := conn.GetOrCreateOwner(ctx, repo.Owner.Login);
		if  err != nil {
			log.Printf("Failed to get or create owner %s: %s", repo.Owner.Login, err)
		}

		dbRepo.Owner = ownerID

		if err := conn.UpsertRepository(ctx, &dbRepo); err != nil {
			log.Printf("Failed to upsert repository %s: %s", repo.FullName, err)
		}
	}
}
