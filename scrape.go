package main

import (
	"fmt"

	"github.com/gocolly/colly"
)

type scrapeData struct {
	Name string
	Price float32
	ReleaseDate string
	ReviewScore string
}

// main() contains code adapted from example found in Colly's docs:
// http://go-colly.org/docs/examples/basic/
func main() {
	// Instantiate default collector
	c := colly.NewCollector()

	// On every a element which has href attribute call callback
	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
                link := e.Attr("href")

		// Print link
                fmt.Printf("Link found: %q -> %s\n", e.Text, link)
	})

	// Before making a request print "Visiting ..."
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	// Start scraping on https://hackerspaces.org
	c.Visit("https://hackerspaces.org/")
}
