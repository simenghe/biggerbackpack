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
	TeamName  string  `json:"team_name"`
	Ranking  string  `json:"ranking"`
	URL float64 `json:"url"`
	Date  string  `json:"date"`
}
// GetTeamData grabs bags from MySQL database
func GetTeamData(){
	db, err := sql.Open("mysql", "simeng:123456@tcp(127.0.0.1:3306)/backpack")
	
}