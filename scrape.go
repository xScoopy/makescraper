package main

import (
	"fmt"
	"time"

	"github.com/gocolly/colly"
)

type gameInfo struct {
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
	c.SetRequestTimeout(120 * time.Second)

	// // On every a element which has href attribute call callback
	// c.OnHTML("a[href]", func(e *colly.HTMLElement) {
    //             link := e.Attr("href")

	// 	// Print link
    //             fmt.Printf("Link found: %q -> %s\n", e.Text, link)
	// })
	
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	c.OnResponse(func(r *colly.Response) {
		fmt.Println("Got a response from", r.Request.URL)
	})

	c.OnError(func(r *colly.Response, e error) {
		fmt.Println("Received error:", e)
	})

	c.Visit("https://store.steampowered.com/search/?filter=topsellers")

}
