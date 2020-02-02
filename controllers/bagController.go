package controllers

import (
	"Backpack/core"
	"encoding/json"
	"fmt"
	"net/http"
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
