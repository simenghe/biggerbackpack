package main

import (
	"Backpack/utils"
	"fmt"
	"strconv"

	"github.com/gocolly/colly"
)

// Item has the fields for a basic amazon catalogue item
type Item struct {
	Name   string  `json:"Name"`
	Desc   string  `json:"Desc"`
	Price  string  `json:"Price"`
	Rating float64 `json:"Rating"`
}

// ScrapeAmazon returns an array of Items from amazon.
func ScrapeAmazon(itemType string) []Item {
	url := fmt.Sprintf("https://www.amazon.ca/s?k=%s&ref=nb_sb_noss_1", itemType)
	var outerQuery = "div.s-result-list.s-search-results.sg-row"
	var innerQuery = "div.a-section.a-spacing-medium"
	var childQuery = "span.a-size-base-plus.a-color-base.a-text-normal"
	var ratingQuery = "span.a-icon-alt"
	var priceQuery = "span.a-price > span.a-offscreen"
	c := colly.NewCollector()
	colly.Async(true)
	fmt.Println(url)
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})
	var itemList []Item
	c.OnHTML(outerQuery, func(e *colly.HTMLElement) {
		e.ForEach(innerQuery, func(_ int, e *colly.HTMLElement) {
			var productName, stars, price string
			productName = e.ChildText(childQuery)
			if productName == "" {
				// If we can't get any name, we return and go directly to the next element
				return
			}
			stars = e.ChildText(ratingQuery)
			rating, err := strconv.ParseFloat(stars, 64)
			if err != nil {
			}
			price = e.ChildText(priceQuery)
			utils.FormatStars(&stars)
			itemList = append(itemList, Item{Name: productName, Desc: "", Price: price, Rating: rating})
			// Create the item to be pushed onto the list.

			// utils.FormatPrice(&price)
			// fmt.Printf("Product Name: %s \nStars: %s \nPrice: %s \n", productName, stars, price)
		})
	})
	c.Visit(url)
	return itemList
}
func main() {
	fmt.Println(ScrapeAmazon("backpack"))
}
