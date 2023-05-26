package main

import (
	"database/sql"
	"log"

	"example.com/test_axxonsoft/v2/config"
	"example.com/test_axxonsoft/v2/database"
	"github.com/gin-gonic/gin"
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

	var postgresConfig = config.PostgresConfig{
		Host:     "localhost",
		Port:     5432,
		Username: "test_axxonsoft",
		Password: "123456",
		Database: "test_axxonsoft",
	}

	var rabbitConfig = config.RabbitMqConfig{
		Host:     "localhost",
		Port:     15672,
		Username: "guest",
		Password: "guest",
	}

	m, err := migrate.New(
		"file://migrations",
		postgresConfig.GetUrl())
	if err != nil {
		log.Fatal(err)
	}
	m.Up()

	// open database
	database.DB, err = sql.Open("postgres", postgresConfig.GetDsn())
	if err != nil {
		panic(err)
	}

	defer database.DB.Close()

	router := gin.Default()

	var appContext = NewApplicationContext(rabbitConfig)

	defer appContext.Close()

	router.GET("/task/:id", appContext.TaskController.GetTask)
	router.POST("/task", appContext.TaskController.PostTask)

	router.Run(":8080")
}
