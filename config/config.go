package config

import "fmt"

type RabbitMqConfig struct {
	Host     string
	Port     int
	Username string
	Password string
}

func (r RabbitMqConfig) getUrl() string {
	return fmt.Sprintf("amqp://%s:%s@%s:%d/", r.Username, r.Password, r.Host, r.Port)
}

type PostgresConfig struct {
	Host     string
	Port     int
	Username string
	Password string
	Database string
}

func (p PostgresConfig) GetUrl() string {
	return "postgres://" + p.Username + ":" + p.Password + "@" + p.Host + ":" + string(p.Port) + "/" + p.Database + "?sslmode=disable"
}

func (p PostgresConfig) GetDsn() string {
	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", p.Host, p.Port, p.Username, p.Password, p.Database)
}
