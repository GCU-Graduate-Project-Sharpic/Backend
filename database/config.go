package database

import (
	"fmt"
	"os"
)

type Config struct {
	host     string
	port     string
	user     string
	password string
	dbName   string
	sslMode  string
}

func NewConfig() *Config {
	return &Config{
		host:     os.Getenv("POSTGRES_HOST"),
		port:     "5432",
		user:     os.Getenv("POSTGRES_USER"),
		password: os.Getenv("POSTGRES_PASSWORD"),
		dbName:   os.Getenv("POSTGRES_DB"),
		sslMode:  "disable",
	}
}

func (c *Config) PsqlConn() string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", c.host, c.port, c.user, c.password, c.dbName, c.sslMode)
}
