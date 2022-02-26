package main

import (
	"fmt"
	"time"

	"github.com/gocolly/colly"
)

type GameInfo struct {
	Name string
	Price float32
	ReleaseDate string
}

// main() contains code adapted from example found in Colly's docs:
// http://go-colly.org/docs/examples/basic/
func main() {
	// Instantiate default collector
	c := colly.NewCollector()
	c.SetRequestTimeout(120 * time.Second)

	// games := make([]GameInfo, 0)

	c.OnHTML("a.search_result_row", func(e *colly.HTMLElement) {
		e.ForEach("div.responsive_search_name_combined", func(i int, h *colly.HTMLElement){
			fmt.Println(h.ChildText("span.title"))
		} )
	})
	
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

