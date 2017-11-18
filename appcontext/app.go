package appcontext

import (
	"crud-app/config"
	sql "database/sql"
	"log"

	_ "github.com/lib/pq"
)

type appContext struct {
	db *sql.DB
}

var context *appContext

func Initiate() {
	db := initDB()
	context = &appContext{
		db: db,
	}
}

func GetDB() *sql.DB {
	return context.db
}

func initDB() *sql.DB {
	var err error
	db, err := sql.Open("postgres", config.DBConfig().ConnectionString())

	if err != nil {
		log.Fatalf("Error connecting to the database")
	}

	if err = db.Ping(); err != nil {
		log.Fatalf("Ping to database host failed: %s \n", err)
	}

	return db
}
