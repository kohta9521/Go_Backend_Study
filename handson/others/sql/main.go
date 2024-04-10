package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)


func main() {
	db, err := sql.Open("sqlite3", "./example.sql")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// table make
	tableName := "user"
	cmd := fmt.Sprintf(`
		CREATE TABLE IF NOT EXISTS %s (
			id SERIAL PRIMART KET NOT NUL,
			name TEXT NOT NULL,
			age INTEGER NOT NULL,
			email TEXT)`, tableName)
		_, err := db.Exec(cmd)
		if err != nil {
			log.Fatal(err)
		}
	)
}