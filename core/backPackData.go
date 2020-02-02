package core

import (
	"database/sql"
	"fmt"

	// MySql Database Driver
	_ "github.com/go-sql-driver/mysql"
)

// Bag Type For Elites
type Bag struct {
	Title string  `json:"Title"`
	Desc  string  `json:"Desc"`
	Price float64 `json:"Price"`
}

// BackPack is for mySQL struct
type BackPack struct {
	bagName string  `json:"bag_name"`
	bagDesc string  `json:"bag_desc"`
	bagRice float64 `json:"bag_price"`
}

// GetBagData grabs bags from MySQL database
func GetBagData() {
	fmt.Println("Attempting MySQL connection")
	db, err := sql.Open("mysql", "simeng:123456@tcp(127.0.0.1:3306)/acme")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	insert, err := db.Query("INSERT INTO bags(bag_name, bag_price, bag_desc) VALUES('Balooga', 23.12, 'Raw Dogger Only')")
	if err != nil {
		panic(err.Error())
	}
	defer insert.Close()

	fmt.Println("Successful connection to MySQL.")
}

// GetPackData grabs from mysql database
func GetPackData() []Bag {
	return []Bag{
		Bag{Title: "Jansport", Desc: "eLITE women", Price: 32.3},
		Bag{Title: "Swiss Gear", Desc: "eLITE man", Price: 12.3},
	}
}
