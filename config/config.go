package config

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

type RabbitMqConfig struct {
	Host     string `mapstructure:"RABBIT_HOST"`
	Port     string `mapstructure:"RABBIT_PORT"`
	Username string `mapstructure:"RABBIT_USERNAME"`
	Password string `mapstructure:"RABBIT_PASSWORD"`
}

func (r RabbitMqConfig) GetUrl() string {
	return fmt.Sprintf("amqp://%s:%s@%s:%s/", r.Username, r.Password, r.Host, r.Port)
}

type PostgresConfig struct {
	Host     string `mapstructure:"PG_HOST"`
	Port     string `mapstructure:"PG_PORT"`
	Username string `mapstructure:"PG_USERNAME"`
	Password string `mapstructure:"PG_PASSWORD"`
	Database string `mapstructure:"PG_DATABASE"`
}

func (p PostgresConfig) GetUrl() string {
	return "postgres://" + p.Username + ":" + p.Password + "@" + p.Host + ":" + string(p.Port) + "/" + p.Database + "?sslmode=disable"
}

func (p PostgresConfig) GetDsn() string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", p.Host, p.Port, p.Username, p.Password, p.Database)
}

var EnvPostgresConfig *PostgresConfig
var EnvRabbitMqConfig *RabbitMqConfig

func InitEnvConfigs() {
	EnvPostgresConfig = loadEnvVariablesForDB()
	EnvRabbitMqConfig = loadEnvVariablesForRabbit()
}

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
