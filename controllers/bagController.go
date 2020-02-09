package controllers

import (
	"Backpack/core"
	"Backpack/scraper"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// Health check for http server
func Health(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Health Works")
}

// FetchTestBags Returns the bags data to the mandems
func FetchTestBags(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Bag Health Test")
	bags := core.GetPackData()
	json.NewEncoder(w).Encode(bags)
}

// FetchBags returns the bag data from the query
func FetchBags(w http.ResponseWriter, r *http.Request) {
	bags := core.GetBagData()
	fmt.Println(bags)
	json.NewEncoder(w).Encode(bags)
}

// AmazonSearch a test function directly scrapes
func AmazonSearch(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	items := scraper.ScrapeAmazon(vars["item"])
	json.NewEncoder(w).Encode(items)
}

// CostcoSearch searches costco.com for stuff
func CostcoSearch(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	scraper.ScrapeCostco(vars["item"])
}
