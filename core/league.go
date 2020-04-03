package core

import (
	"database/sql"
	"fmt"

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

	// perform an insert operation
	insert, err := db.Query(`
	INSERT INTO players
	(roster_name,roster_id,player_name,player_role,created_date)
  	VALUES
	('SimpSquad', '123456', 'Bobby', 'ADC', now());
	`)
	if err != nil {
		panic(err.Error())
	}
	defer insert.Close()

	fmt.Println("Successful connection to MySQL.")
}

// GetRoster() gets the roster with