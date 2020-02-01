package controllers

import(
	"fmt"
	"net/http"
	"encoding/json"
)
// Bag Type For Elites
type Bag struct {
	Title string  `json:"Title"`
	Desc  string  `json:"Desc"`
	Price float64 `json:"Price"`
}

// FetchBags Returns the bags data to the mandems
func FetchBags(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Fetching Bags")
	bags := []Bag{
		Bag{Title: "Jansport", Desc: "eLITE women", Price: 32.3},
		Bag{Title: "Swiss Gear", Desc: "eLITE man", Price: 12.3},
	}
	json.NewEncoder(w).Encode(bags)
}