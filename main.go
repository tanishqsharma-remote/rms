package main

import (
	"RMS/database_dir"
	"RMS/router_dir"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
	"log"
	"net/http"
)

func main() {

	db := database_dir.DBconnect()
	driver, dbErr := postgres.WithInstance(db, &postgres.Config{})
	if dbErr != nil {
		log.Fatal(dbErr)
	}

	m, err := migrate.NewWithDatabaseInstance("file://database_dir/migrations_dir", "postgres", driver)
	if err != nil {
		log.Fatal(err)
	}
	er := m.Up()
	if er == migrate.ErrNoChange {
		//
	}
	r := router_dir.Router()
	err1 := http.ListenAndServe(":8080", r)
	if err1 != nil {
		log.Fatal("Error")
		return
	}
}
