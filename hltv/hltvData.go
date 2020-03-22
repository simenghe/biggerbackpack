package hltv

import (
	"database/sql"
	"fmt"

	// "encoding/json"
	// MySql Database Driver
	_ "github.com/go-sql-driver/mysql"
)

// CSGOteam main bag
type CSGOteam struct {
	ID       int     `json:"ID"`
	TeamName string  `json:"team_name"`
	Ranking  string  `json:"ranking"`
	URL      string `json:"url"`
	Date     string  `json:"date"`
}

// GetTeamData grabs bags from MySQL database
func GetTeamData() []CSGOteam {
	db, err := sql.Open("mysql", "docker:docker@tcp(127.0.0.1:3308)/backpack")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	results, err := db.Query("SELECT * FROM csgoteams")
	if err != nil {
		panic(err.Error())
	}
	csgoTeams := []CSGOteam{}
	for results.Next() {
		var team CSGOteam
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
