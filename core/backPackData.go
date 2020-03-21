package core

import (
	"database/sql"
	"fmt"
	// "encoding/json"
	// MySql Database Driver
	_ "github.com/go-sql-driver/mysql"
)

// Bag2 Type For Elites
type Bag2 struct {
	Title string  `json:"Title"`
	Desc  string  `json:"Desc"`
	Price float64 `json:"Price"`
}

// Bag main bag
type Bag struct {
	ID       int     `json:"ID"`
	BagName  string  `json:"bag_name"`
	BagDesc  string  `json:"bag_desc"`
	BagPrice float64 `json:"bag_price"`
}

// GetBagData grabs bags from MySQL database
func GetBagData() []Bag {
	fmt.Println("Attempting MySQL connection")
	db, err := sql.Open("mysql", "simeng:123456@tcp(127.0.0.1:3306)/acme")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	results, err := db.Query("SELECT * FROM bags")
	if err != nil {
		panic(err.Error())
	}
	bags := []Bag{}
	for results.Next() {
		var bagee Bag
		err = results.Scan(&bagee.ID, &bagee.BagName, &bagee.BagPrice, &bagee.BagDesc)
		if err != nil {
			panic(err.Error())
		}
		bags = append(bags, bagee)
	}
	fmt.Println("Successful connection to MySQL.")
	return bags
}

// GetPackData grabs from mysql database
func GetPackData() []Bag2 {
	return []Bag2{
		Bag2{Title: "Jansport", Desc: "eLITE women", Price: 32.3},
		Bag2{Title: "Swiss Gear", Desc: "eLITE man", Price: 12.3},
	}
}
