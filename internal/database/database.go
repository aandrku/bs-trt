package database

import (
	"bufio"
	"database/sql"
	"log"
	"os"
	_ "github.com/go-sql-driver/mysql"
)

var dsn string

func init() {
	file, err := os.Open("./secret/dsn.database")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)

	if scanner.Scan() {
		dsn = scanner.Text()
		return	
	}
}

func ConnectionDB() *sql.DB {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}
	
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	return db
}
