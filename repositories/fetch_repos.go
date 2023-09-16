package repositories

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/letsfixoss/gh-sponsor-grabber/internal"
)

const perPage = 100
const maxPages = 10
const timeout = 30 * time.Second

// Repository structure returned by GitHub API
type Repository struct {
	FullName string `json:"full_name"`
}

// SearchResult to hold the array of repositories
type SearchResult struct {
	Items []Repository `json:"items"`
}

func GetRepos() []Repository {
	client := &http.Client{Timeout: timeout}
	env := internal.GetEnv()
	token := env.GithubToken
	foundRepos := make([]Repository, 0)

	// Loop over pages
	for page := 1; page <= maxPages; page++ {
		url := fmt.Sprintf("https://api.github.com/search/repositories?q=language:go&sort=stars&order=desc&per_page=%d&page=%d", perPage, page)
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			fmt.Println("Error creating HTTP request:", err)
			return nil
		}

		req.Header.Set("Authorization", "token "+token)
		resp, err := client.Do(req)
		if err != nil {
			fmt.Println("Error making HTTP request:", err)
			return nil
		}

		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("Error reading HTTP body:", err)
			return nil
		}

		var result SearchResult
		err = json.Unmarshal(body, &result)
		if err != nil {
			fmt.Println("Error unmarshaling JSON:", err)
			return nil
		}

		// Loop over each repository
		for _, repo := range result.Items {
			// Download the dependency files here
			// For example: Fetch package.json or go.mod from the repo
			fmt.Println("Fetching files for repo:", repo.FullName)
			// Implement code to fetch package.json, go.mod, etc.
			foundRepos = append(foundRepos, repo)
		}

		// Sleep between requests to respect rate limits
		time.Sleep(1 * time.Second)
	}

	return foundRepos
}
