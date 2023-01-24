package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"example.com/test_axxonsoft/v2/controller"
	"example.com/test_axxonsoft/v2/database"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "test_axxonsoft"
	password = "123456"
	dbname   = "test_axxonsoft"
)

func main() {

	m, err := migrate.New(
		"file://migrations",
		"postgres://test_axxonsoft:123456@localhost:5432/test_axxonsoft?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	m.Up()

	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	// open database
	database.DB, err = sql.Open("postgres", psqlconn)
	if err != nil {
		panic(err)
	}

	http.HandleFunc("/task/", controller.TaskHandler)

	http.ListenAndServe(":8080", nil)
}
