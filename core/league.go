package core

import (
	"fmt"
	"database/sql"
	// "encoding/json"
	// MySql Database Driver
	_ "github.com/go-sql-driver/mysql"
)

// CreateRoster saves a 5 man roster
func CreateRoster() {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3308)/backpack")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	if err != nil {
		panic(err.Error())
	}
	// csgoTeams := []hltvScraper.CSGOteam{}
	// for results.Next() {
	// 	var team hltvScraper.CSGOteam
	// 	err = results.Scan(&team.ID, &team.TeamName, &team.Ranking, &team.URL, &team.Date)
	// 	if err != nil {
	// 		panic(err.Error())
	// 	}
	// 	csgoTeams = append(csgoTeams, team)
	// }
	fmt.Println("Successful connection to MySQL.")
}
