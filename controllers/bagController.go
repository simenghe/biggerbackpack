package controllers

import (
	"Backpack/core"
	"encoding/json"
	"fmt"
	"net/http"
)

// FetchBags Returns the bags data to the mandems
func FetchBags(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Bag Health Test")
	bags := core.GetPackData()
	json.NewEncoder(w).Encode(bags)
}