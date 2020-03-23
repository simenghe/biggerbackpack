package hltv

import (
	"database/sql"
	"fmt"
	hltvScraper "Backpack/scraper"
	// "encoding/json"
	// MySql Database Driver
	_ "github.com/go-sql-driver/mysql"
)

// GetTeamData grabs bags from MySQL database
func GetTeamData() []hltvScraper.CSGOteam {
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3308)/backpack")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	results, err := db.Query("SELECT * FROM csgoteams")
	if err != nil {
		panic(err.Error())
	}
	csgoTeams := []hltvScraper.CSGOteam{}
	for results.Next() {
		var team hltvScraper.CSGOteam
		err = results.Scan(&team.ID, &team.TeamName, &team.Ranking, &team.URL, &team.Date)
		if err != nil {
			panic(err.Error())
		}
		csgoTeams = append(csgoTeams, team)
	}
	fmt.Println("Successful connection to MySQL.")
	fmt.Println(csgoTeams)
	return csgoTeams
}
