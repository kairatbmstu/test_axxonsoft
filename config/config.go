package config

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

/*
The RabbitMqConfig struct represents the configuration variables related to RabbitMQ.
It has fields for the RabbitMQ host, port, username, and password.
The mapstructure tags are used for mapping the environment variables to struct fields.
*/
type RabbitMqConfig struct {
	Host     string `mapstructure:"RABBIT_HOST"`
	Port     string `mapstructure:"RABBIT_PORT"`
	Username string `mapstructure:"RABBIT_USERNAME"`
	Password string `mapstructure:"RABBIT_PASSWORD"`
}

/*
The GetUrl method of the RabbitMqConfig struct returns the RabbitMQ connection
URL string based on the configuration variables. It formats the URL using the username, password, host, and port.
*/
func (r RabbitMqConfig) GetUrl() string {
	return fmt.Sprintf("amqp://%s:%s@%s:%s/", r.Username, r.Password, r.Host, r.Port)
}

/*
The PostgresConfig struct represents the configuration variables related to PostgreSQL.
It has fields for the PostgreSQL host, port, username, password, and database name.
The mapstructure tags are used for mapping the environment variables to struct fields.
*/
type PostgresConfig struct {
	Host     string `mapstructure:"PG_HOST"`
	Port     string `mapstructure:"PG_PORT"`
	Username string `mapstructure:"PG_USERNAME"`
	Password string `mapstructure:"PG_PASSWORD"`
	Database string `mapstructure:"PG_DATABASE"`
}

/*
The GetUrl method of the PostgresConfig struct returns the PostgreSQL
connection URL string based on the configuration variables.
It constructs the URL by concatenating the username, password, host, port, database name, and SSL mode parameters.
*/
func (p PostgresConfig) GetUrl() string {
	return "postgres://" + p.Username + ":" + p.Password + "@" + p.Host + ":" + string(p.Port) + "/" + p.Database + "?sslmode=disable"
}

/*
The GetDsn method of the PostgresConfig struct returns the PostgreSQL
connection DSN (Data Source Name) string based on the configuration variables.
It formats the DSN using the host, port, username, password, and database name.
*/
func (p PostgresConfig) GetDsn() string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", p.Host, p.Port, p.Username, p.Password, p.Database)
}

/*
The EnvPostgresConfig variable holds the PostgreSQL configuration loaded from the environment variables.
*/
var EnvPostgresConfig *PostgresConfig

/*
The EnvRabbitMqConfig variable holds the RabbitMQ configuration loaded from the environment variables.
*/
var EnvRabbitMqConfig *RabbitMqConfig

/*
The InitEnvConfigs function is used to initialize the environment
configurations by loading the environment variables from the "app.env" file.
It reads the configuration file using viper and unmarshals the values into
the corresponding structs (EnvPostgresConfig and EnvRabbitMqConfig).
*/
func InitEnvConfigs() {
	EnvPostgresConfig = loadEnvVariablesForDB()
	EnvRabbitMqConfig = loadEnvVariablesForRabbit()
}

/*
The loadEnvVariablesForDB function loads the environment variables related to
the PostgreSQL configuration from the "app.env" file using viper.
It unmarshals the values into the PostgresConfig struct and returns the populated configuration.
*/
func loadEnvVariablesForDB() (config *PostgresConfig) {
	viper.AddConfigPath(".")
	viper.SetConfigName("app")
	viper.SetConfigType("env")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatal("Error reading env file", err)
	}

	if err := viper.Unmarshal(&config); err != nil {
		log.Fatal(err)
	}
	return
}

/*
The loadEnvVariablesForRabbit function loads the environment variables
related to the RabbitMQ configuration from the "app.env" file using viper.
It unmarshals the values into the RabbitMqConfig struct and returns the populated configuration.
*/
func loadEnvVariablesForRabbit() (config *RabbitMqConfig) {
	viper.AddConfigPath(".")
	viper.SetConfigName("app")
	viper.SetConfigType("env")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatal("Error reading env file", err)
	}

	if err := viper.Unmarshal(&config); err != nil {
		log.Fatal(err)
	}
	return
}
