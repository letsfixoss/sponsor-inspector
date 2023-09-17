package main

import (
	"context"
	"fmt"
	"log"
	"strconv"

	colly "github.com/gocolly/colly/v2"
	"github.com/letsfixoss/gh-sponsor-grabber/db"
)

const sponsorCountSelector = "#sponsors-section-list > div:nth-child(1) > h4 > span"

func main() {
	ctx := context.Background()
	conn := db.GetConnection()

	fmt.Printf("Getting repository owners...\n")

	ro, err := conn.GetRepoOwners(ctx)
	if err != nil {
		log.Fatalf("failed to get repo owners: %s", err)
	}

	fmt.Printf("Found %d repo owners\n", len(ro))

	for _, r := range ro {
		fmt.Printf("Getting sponsor count for %s...\n", r.Name)

		count, err := scrapeSponsors(r)
		if err != nil {
			log.Fatalf("failed to scrape sponsors: %s", err)
		}

		if count == nil {
			fmt.Printf("No sponsors found for %s\n", r.Name)
		} else {
			fmt.Printf("Found %v sponsors for %s\n", *count, r.Name)
		}
		fmt.Printf("Updating sponsor count for %s...\n", r.Name)

		if err := conn.UpdateSponsorCount(ctx, r.ID, count); err != nil {
			log.Fatalf("failed to update sponsor count: %s", err)
		}
	}
}

func scrapeSponsors(ro *db.RepoOwners) (*uint, error) {
	var count *uint
	var retErr error
	url := fmt.Sprintf("https://github.com/sponsors/%s", ro.Name)
	c := colly.NewCollector()

	c.OnHTML(sponsorCountSelector, func(e *colly.HTMLElement) {
		foundCount, err := strconv.Atoi(e.Text)
		if err != nil {
			retErr = fmt.Errorf("failed to convert sponsor count to int: %s", err)
			return
		}

		foundCountUint := uint(foundCount)
		count = &foundCountUint
	})

	c.Visit(url)

	return count, retErr
}
