package main

import (
	"log"
	"os"
	"testing"
	"webapp/pkg/db"
)

var app application

func TestMain(m *testing.M) {
	pathToTemplates = "./../../templates"

	app.Session = getSession()
	app.DSN = "host=localhost port=5432 user=postgres password=postgres sslmode=disable dbname=users timezone=UTC connect_timeout=5"

	// connect to the database
	conn, err := app.connectToDB()
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	app.DB = db.PostgresConn{DB: conn}

	os.Exit(m.Run())

}
