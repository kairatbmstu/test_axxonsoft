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

func main() {
	config.InitEnvConfigs()

	m, err := migrate.New(
		"file://migrations",
		config.EnvPostgresConfig.GetUrl())
	if err != nil {
		log.Fatal(err)
	}
	m.Up()

	// open database
	database.DB, err = sql.Open("postgres", config.EnvPostgresConfig.GetDsn())
	if err != nil {
		panic(err)
	}

	defer database.DB.Close()

	router := gin.Default()

	var appContext = NewApplicationContext(*config.EnvRabbitMqConfig)

	defer appContext.Close()

	router.GET("/task/:id", appContext.TaskController.GetTask)
	router.POST("/task", appContext.TaskController.PostTask)

	router.Run(":8080")
}
