package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Bag struct {
	Title string  `json:"Title"`
	Desc  string  `json:"Desc"`
	Price float64 `json:"Price"`
}

var Bags []Bag

func fetchBags(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Fetching Bags")
	bags := []Bag{
		Bag{Title: "Jansport", Desc: "eLITE women", Price: 32.3},
		Bag{Title: "Swiss Gear", Desc: "eLITE man", Price: 12.3},
	}
	json.NewEncoder(w).Encode(bags)
}
func handleReq() {
	fmt.Println("Handling the requests!")
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", fetchBags)
	log.Fatal(http.ListenAndServe(":10000", myRouter))
}

func main() {
	fmt.Println("Main is running")
	handleReq()
}
