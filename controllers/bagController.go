package controllers

import(
	"fmt"
	"net/http"
	"encoding/json"
	"Backpack/core"
)
// Bag Type For Elites
// type Bag struct {
// 	Title string  `json:"Title"`
// 	Desc  string  `json:"Desc"`
// 	Price float64 `json:"Price"`
// }

// FetchBags Returns the bags data to the mandems
func FetchBags(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Bag Health Test")
	bags := core.GetPackData()
	json.NewEncoder(w).Encode(bags)
}