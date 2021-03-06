package controllers

import (
	"Backpack/core"
	"encoding/json"

	// "fmt"
	"net/http"
	// "github.com/gorilla/mux"
)

// LeagueRoster teams grabs the stored teams from the database.
func LeagueRoster(w http.ResponseWriter, r *http.Request) {
	var player core.Player
	_ = json.NewDecoder(r.Body).Decode(&player)
	core.CreateRoster(player)
	// json.NewEncoder(w).Encode(teams)
}

// // FetchTestBags Returns the bags data to the mandems
// func FetchTestBags(w http.ResponseWriter, r *http.Request) {
// 	fmt.Println("Bag Health Test")
// 	bags := core.GetPackData()
// 	json.NewEncoder(w).Encode(bags)
// }
