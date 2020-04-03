package core

import (
	"database/sql"
	"fmt"

	// "encoding/json"
	// MySql Database Driver
	_ "github.com/go-sql-driver/mysql"
)

// Player struct defines basic league player
type Player struct {
	PlayerName string `json:"playerName"`
	PlayerRole string `json:"playerRole"`
	RosterID   string `json:"rosterID"`
	RosterName string `json:"rosterName"`
}

// CreateRoster saves a 5 man roster
func CreateRoster(player Player) {
	db, err := sql.Open("mysql", "docker:docker@tcp(localhost:3308)/backpack")
	if err != nil {
		fmt.Println("loss")
		panic(err.Error())
	}
	defer db.Close()
	sqlQuery := fmt.Sprintf(`
	INSERT INTO players
	(roster_name,roster_id,player_name,player_role,created_date)
  	VALUES
	('%s', '%s', '%s', '%s', now());
	`, player.RosterName, player.RosterID, player.PlayerName, player.PlayerRole)
	// perform an insert operation
	insert, err := db.Query(sqlQuery)
	if err != nil {
		panic(err.Error())
	}
	defer insert.Close()

	fmt.Println("Successful connection to MySQL.")
}

// GetRoster() gets the roster with
