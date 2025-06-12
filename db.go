package main

import (
	"database/sql"
	"io/ioutil"
	"log"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func initDB() {
	var err error

	// Step 1: Connect without specifying the DB (to ensure creation)
	db, err = sql.Open("mysql", "Students:rana@tcp(127.0.0.1:3306)/")
	if err != nil {
		log.Fatal("DB connection error:", err)
	}

	// Step 2: Load and execute tables.sql
	content, err := ioutil.ReadFile("tables.sql")
	if err != nil {
		log.Fatal("Could not read tables.sql:", err)
	}

	commands := strings.Split(string(content), ";")
	for _, cmd := range commands {
		cmd = strings.TrimSpace(cmd)
		if cmd != "" {
			if _, err := db.Exec(cmd); err != nil {
				log.Fatalf("Error executing SQL command: %s\n%v", cmd, err)
			}
		}
	}

	// Step 3: Reconnect using the actual DB
	db, err = sql.Open("mysql", "Students:rana@tcp(127.0.0.1:3306)/myproject")
	if err != nil {
		log.Fatal("DB reconnect error:", err)
	}

	if err = db.Ping(); err != nil {
		log.Fatal("DB ping failed:", err)
	}

	log.Println("Connected to database and initialized")
}
