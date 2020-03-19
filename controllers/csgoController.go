package controllers

import (
	// "Backpack/core"
	"Backpack/hltv"
	"encoding/json"

	// "fmt"
	"net/http"
	// "github.com/gorilla/mux"
)

// CSGOteams grabs the stored teams from the database.
func CSGOteams(w http.ResponseWriter, r *http.Request) {
	teams := hltv.GetTeamData()
	json.NewEncoder(w).Encode(teams)
}

// // FetchTestBags Returns the bags data to the mandems
// func FetchTestBags(w http.ResponseWriter, r *http.Request) {
// 	fmt.Println("Bag Health Test")
// 	bags := core.GetPackData()
// 	json.NewEncoder(w).Encode(bags)
// }
