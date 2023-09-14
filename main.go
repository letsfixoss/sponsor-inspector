package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
)

// Repository structure returned by GitHub API
type Repository struct {
	FullName string `json:"full_name"`
}

// SearchResult to hold the array of repositories
type SearchResult struct {
	Items []Repository `json:"items"`
}

func main() {
	godotenv.Load()
	client := &http.Client{
		Timeout: time.Second * 30,
	}

	token := os.Getenv("GITHUB_TOKEN")

	// Loop over pages
	for page := 1; page <= 10; page++ {
		url := fmt.Sprintf("https://api.github.com/search/repositories?q=stars:>1000&sort=stars&order=desc&page=%d&per_page=100", page)
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			fmt.Println("Error creating HTTP request:", err)
			return
		}

		req.Header.Set("Authorization", "token "+token)
		resp, err := client.Do(req)
		if err != nil {
			fmt.Println("Error making HTTP request:", err)
			return
		}

		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("Error reading HTTP body:", err)
			return
		}

		var result SearchResult
		err = json.Unmarshal(body, &result)
		if err != nil {
			fmt.Println("Error unmarshaling JSON:", err)
			return
		}

		// Loop over each repository
		for _, repo := range result.Items {
			// Download the dependency files here
			// For example: Fetch package.json or go.mod from the repo
			fmt.Println("Fetching files for repo:", repo.FullName)
			// Implement code to fetch package.json, go.mod, etc.
		}

		// Sleep between requests to respect rate limits
		time.Sleep(1 * time.Second)
	}
}
