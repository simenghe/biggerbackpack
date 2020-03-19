package routes

import (
	bagController "Backpack/controllers"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// HandleReq handles our main route
func HandleReq() {
	fmt.Println("Handling the requests!")
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", bagController.FetchTestBags)
	myRouter.HandleFunc("/health", bagController.Health)
	myRouter.HandleFunc("/getbags", bagController.FetchBags)

	// Scraper test route, cannot rely on this for data.
	myRouter.HandleFunc("/amazon/{item}", bagController.AmazonSearch)
	myRouter.HandleFunc("/costco/{item}", bagController.CostcoSearch)

	// Hltv /csgo stuff
	myRouter.HandleFunc("/csgoteams", bagController.CSGOteams)
	log.Fatal(http.ListenAndServe(":5000", myRouter))
}
