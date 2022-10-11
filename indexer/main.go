package main

import (
	"fmt"
	"log"
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "./index.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	rows, err := db.Query("select name from erc20")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		var name string
		err = rows.Scan(&name)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("ERC20 Name: ", name)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
}
