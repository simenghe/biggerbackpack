package scraper

import (
	"fmt"
	"log"
	"regexp"

	"github.com/PuerkitoBio/goquery"
)

// StockScrape Scrapes Yahoo Finance For stock prices.
func StockScrape() {
	url := "https://finance.yahoo.com/quote/FB?p=FB&.tsrc=fin-tre-srch"
	doc, err := goquery.NewDocument(url)
	if err != nil {
		log.Fatal(err)
	}

	var price string
	fmt.Println(doc.Find("div#quote-header-info").Length())
	// var caneText string
	doc.Find("div#quote-header-info").Each(func(i int, s *goquery.Selection) {
		price = s.Find("span").Text()
		found, _ := regexp.MatchString("USD([0-9]|.+)-", price)
		fmt.Println(found)
		r, _ := regexp.Compile("USD([0-9|.]*)")
		fmt.Println(r.FindString(price))
	})
}
