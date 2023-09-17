package main

import (
	colly "github.com/gocolly/colly/v2"
)

const sponsorCountSelector = "#sponsors-section-list > div:nth-child(1) > h4 > span"

func main() {
	c := colly.NewCollector()
	c.OnHTML(sponsorCountSelector, func(e *colly.HTMLElement) {
		println(e.Text)
	})

	c.Visit("https://github.com/sponsors/bigskysoftware")
}
