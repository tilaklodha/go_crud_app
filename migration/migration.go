package migration

import (
	"fmt"
	"go_crud_app/config"
	"log"
	"strings"

	_ "github.com/mattes/migrate/driver/postgres"
	"github.com/mattes/migrate/migrate"
	piped "github.com/mattes/migrate/pipe"
)

const dbMigrationsPath = "./migration/queries"

func RunDataBaseMigrations() error {
	connectionURL := config.DBConfig().ConnectionURL()
	log.Println(connectionURL)
	allErrors, ok := migrate.UpSync(connectionURL, dbMigrationsPath)
	if !ok {
		return joinErrors(allErrors)
	}
	log.Println("Migration Successful")
	return nil
}

func RollBackLatestMigration() error {
	pipe := piped.New()
	go migrate.Migrate(pipe, config.DBConfig().ConnectionURL(), dbMigrationsPath, -1)
	return joinErrors(piped.ReadErrors(pipe))
}

func joinErrors(errors []error) error {
	var errorMsgs []string
	for _, err := range errors {
		errorMsgs = append(errorMsgs, err.Error())
	}
	return fmt.Errorf(strings.Join(errorMsgs, ","))
}
