package scraper

import (
	// "Backpack/utils"
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

	// main cane
	c := colly.NewCollector()
	var header = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_9_3) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/35.0.1916.47 Safari/537.36"
	colly.Async(true)
	fmt.Println(url)
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
		r.Headers.Set("User-Agent", header)

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
			// utils.FormatStars(&stars)
			itemList = append(itemList, Item{Name: productName, Desc: "", Price: price, Rating: rating})
			// Create the item to be pushed onto the list.

			// utils.FormatPrice(&price)
			// fmt.Printf("Product Name: %s \nStars: %s \nPrice: %s \n", productName, stars, price)
		})
	})
	c.Visit(url)
	fmt.Println(itemList)
	return itemList
}

// ScrapeCostco scrapes items from costco.com
func ScrapeCostco(itemType string) {
	url := fmt.Sprintf("https://www.costco.com/CatalogSearch?dept=All&keyword=%s/", itemType)
	// main cane
	c := colly.NewCollector()
	colly.Async(true)
	fmt.Println(url)
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})
	
	c.OnHTML("", func(e *colly.HTMLElement){
		// e.Request.Visit(e.Attr("href"))
		fmt.Println(e.Text)
	})
	c.Visit(url)
}
