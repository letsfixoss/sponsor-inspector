package main

import (
	"fmt"

	"github.com/joho/godotenv"
	"github.com/letsfixoss/gh-sponsor-grabber/repositories"
)

func main() {
	godotenv.Load()

	repos := repositories.GetRepos()

	fmt.Printf("found %d repos\n", len(repos))
}
