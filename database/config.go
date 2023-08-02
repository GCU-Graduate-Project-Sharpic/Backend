package database

import (
	"fmt"
	"os"
)

type Config struct {
	PsqlConfig  *PsqlConfig
	MinioConfig *MinioConfig
}

type MinioConfig struct {
	Host     string
	AccessID string
	AccessPW string
	useSSL   bool
}

type PsqlConfig struct {
	host     string
	port     string
	user     string
	password string
	dbName   string
	sslMode  string
}

func NewConfig() *Config {
	return &Config{
		PsqlConfig: &PsqlConfig{
			host:     os.Getenv("POSTGRES_HOST"),
			port:     "5432",
			user:     os.Getenv("POSTGRES_USER"),
			password: os.Getenv("POSTGRES_PASSWORD"),
			dbName:   os.Getenv("POSTGRES_DB"),
			sslMode:  "disable",
		},
		MinioConfig: &MinioConfig{
			Host:     os.Getenv("MINIO_HOST"),
			AccessID: os.Getenv("MINIO_ACCESS_ID"),
			AccessPW: os.Getenv("MINIO_ACCESS_PW"),
			useSSL:   false,
		},
	}
}

func (c *Config) PsqlConn() string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", c.PsqlConfig.host, c.PsqlConfig.port, c.PsqlConfig.user, c.PsqlConfig.password, c.PsqlConfig.dbName, c.PsqlConfig.sslMode)
}
