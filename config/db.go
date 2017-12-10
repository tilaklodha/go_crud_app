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
		name:     "go_crud_app_test",
		username: "postgres",
		password: "",
	}
}

func (dc *dbConfig) ConnectionString() string {
	return fmt.Sprintf("dbname=%s user=%s password='%s' host=%s sslmode=disable", dc.name, dc.username, dc.password, dc.host)
}

func (dc *dbConfig) ConnectionURL() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable", dc.username, dc.password, dc.host, dc.port, dc.name)
}
