package main

import (
	"Backpack/utils"
	"fmt"

	"github.com/gocolly/colly"
)

// Item has the fields for a basic amazon catalogue item
type Item struct {
	Name   string  `json:"Name"`
	Desc   string  `json:"Desc"`
	Price  float64 `json:"Price"`
	Rating float32 `json:"Rating"`
}

// ScrapeAmazon returns an array of Items from amazon.
func ScrapeAmazon() []Item {
	return nil
}
func main() {
	c := colly.NewCollector()

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	c.OnHTML("div.s-result-list.s-search-results.sg-row", func(e *colly.HTMLElement) {
		e.ForEach("div.a-section.a-spacing-medium", func(_ int, e *colly.HTMLElement) {
			var productName, stars, price string
			productName = e.ChildText("span.a-size-base-plus.a-color-base.a-text-normal")
			if productName == "" {
				// If we can't get any name, we return and go directly to the next element
				return
			}
			stars = e.ChildText("span.a-icon-alt")
			price = e.ChildText("span.a-price > span.a-offscreen")
			utils.FormatStars(&stars)
			// utils.FormatPrice(&price)
			fmt.Printf("Product Name: %s \nStars: %s \nPrice: %s \n", productName, stars, price)
		})
	})

	c.Visit("https://www.amazon.ca/s?k=backpack&ref=nb_sb_noss_1")
}
