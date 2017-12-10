package appcontext

import (
	"go_crud_app/config"
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type appContext struct {
	db *sqlx.DB
}

var context *appContext

func Initiate() {
	db := initDB()
	context = &appContext{
		db: db,
	}
}

func GetDB() *sqlx.DB {
	return context.db
}

func initDB() *sqlx.DB {
	var err error
	db, err := sqlx.Open("postgres", config.DBConfig().ConnectionString())

	if err != nil {
		log.Fatalf("Error connecting to the database")
	}

	if err = db.Ping(); err != nil {
		log.Fatalf("Ping to database host failed: %s \n", err)
	}

	return db
}
