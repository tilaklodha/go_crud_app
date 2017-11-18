package config

import "fmt"

type dbConfig struct {
	host     string
	port     int
	username string
	password string
	name     string
}

func newDBConfig() *dbConfig {
	return &dbConfig{
		host:     "localhost",
		port:     5432,
		name:     "users",
		username: "postgres",
		password: "",
	}
}

func (dc *dbConfig) ConnectionString() string {
	return fmt.Sprintf("dbname=%s user=%s password='%s' host=%s sslmode=disable", dc.name, dc.username, dc.password, dc.host)
}
