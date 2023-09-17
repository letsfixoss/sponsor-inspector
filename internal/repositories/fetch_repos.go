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
	FullName        string `json:"full_name"`
	StargazersCount int    `json:"stargazers_count"`
	AvatarURL       string `json:"avatar_url"`
	ForksCount      int    `json:"forks_count"`
	WatchersCount   int    `json:"watchers_count"`
	Language        string `json:"language"`
	CreatedAt       string `json:"created_at"`
	ID              uint64  `json:"id"`
	URL             string `json:"url"`
	OpenIssuesCount int    `json:"open_issues_count"`
	Archived        bool   `json:"archived"`
	Disabled        bool   `json:"disabled"`
	Owner           struct {
		Login     string `json:"login"`
		AvatarURL string `json:"avatar_url"`
		ID        uint64 `json:"id"`
		URL       string `json:"url"`
	} `json:"owner"`
}

func (r Repository) CreatedAsTime() time.Time {
	t, err := time.Parse(time.RFC3339, r.CreatedAt)
	if err != nil {
		fmt.Println("Error parsing time:", err)
		return time.Now() // well this could be bad...
	}

	return t
}

// SearchResult to hold the array of repositories
type SearchResult struct {
	Items []Repository `json:"items"`
}

// GetRepos returns a list of repositories
// @see https://docs.github.com/en/free-pro-team@latest/rest/search/search?apiVersion=2022-11-28#search-repositories
func GetRepos() []Repository {
	client := &http.Client{Timeout: timeout}
	env := internal.GetEnv()
	token := env.GithubToken
	foundRepos := make([]Repository, 0)

	// Loop over pages
	for page := 1; page <= maxPages; page++ {
		url := fmt.Sprintf("https://api.github.com/search/repositories?q=stars:>10000&sort=stars&order=desc&per_page=%d&page=%d", perPage, page)
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
			foundRepos = append(foundRepos, repo)
		}

		// Sleep between requests to respect rate limits
		time.Sleep(1 * time.Second)
	}

	return foundRepos
}
