package hltv

import (
	hltvScraper "Backpack/scraper"
	"fmt"

	// "encoding/json"
	// MySql Database Driver
	_ "github.com/go-sql-driver/mysql"
)

// GetTeamData grabs bags from MySQL database
func GetTeamData() []hltvScraper.CSGOteam {
	// db, err := sql.Open("mysql", "root:root@tcp(localhost:3308)/backpack")
	// if err != nil {
	// 	panic(err.Error())
	// }
	// defer db.Close()
	// results, err := db.Query("SELECT * FROM csgoteams")
	// if err != nil {
	// 	panic(err.Error())
	// }
	// csgoTeams := []hltvScraper.CSGOteam{}
	// for results.Next() {
	// 	var team hltvScraper.CSGOteam
	// 	err = results.Scan(&team.ID, &team.TeamName, &team.Ranking, &team.URL, &team.Date)
	// 	if err != nil {
	// 		panic(err.Error())
	// 	}
	// 	csgoTeams = append(csgoTeams, team)
	// }
	// fmt.Println("Successful connection to MySQL.")
	csgoTeams := hltvScraper.ScrapeHltvTeams()
	fmt.Println(csgoTeams)
	return csgoTeams
}
