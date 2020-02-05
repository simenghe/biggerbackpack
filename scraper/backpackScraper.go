package main

import (
	"fmt"
	"regexp"

	"github.com/gocolly/colly"
)

func FormatPrice(price *string) {
	r := regexp.MustCompile(`\$(\d+(\.\d+)?).*$`)

	newPrices := r.FindStringSubmatch(*price)

	if len(newPrices) > 1 {
		*price = newPrices[1]
	} else {
		*price = "Unknown"
	}

}

func FormatStars(stars *string) {
	if len(*stars) >= 3 {
		*stars = (*stars)[0:3]
	} else {
		*stars = "Unknown"
	}
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
			FormatStars(&stars)
			// FormatPrice(&price)
			fmt.Printf("Product Name: %s \nStars: %s \nPrice: %s \n", productName, stars, price)
		})
	})

	c.Visit("https://www.amazon.ca/s?k=backpack&ref=nb_sb_noss_2")
}
