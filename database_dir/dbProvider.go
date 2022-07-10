package database_dir

import (
	//"database/sql"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"log"
)

func DbConnect() *sqlx.DB {
	connStr := "user=postgres dbname=postgres password= postgres sslmode=disable"
	database, err := sqlx.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	return database
}
