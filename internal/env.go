package internal

import (
	"os"

	"github.com/joho/godotenv"
)

type Env struct {
	GithubToken string
}

func init() {
	godotenv.Load()

	if os.Getenv("GITHUB_TOKEN") == "" {
		panic("GITHUB_TOKEN is not set")
	}
}

func GetEnv() *Env {
	return &Env{
		GithubToken: os.Getenv("GITHUB_TOKEN"),
	}
}
