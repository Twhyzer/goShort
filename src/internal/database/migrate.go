package database

import (
	"database/sql"
	"embed"
	"regexp"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/golang-migrate/migrate/v4/source/iofs"
	"github.com/spf13/cast"
)

//go:embed structure/*.sql
var dbStructureFS embed.FS

func MigrateDatabase(db *sql.DB, databaseName string) {
	err := migrateFromEmbedFS(db, dbStructureFS, databaseName)
	if err != nil {
		panic(err)
	}
}

func migrateFromEmbedFS(db *sql.DB, embedFS embed.FS, databaseName string) error {

	fileContent, fileContentErr := iofs.New(embedFS, "structure")
	if fileContentErr != nil {
		return fileContentErr
	}

	postgresDriver, postgresDriverErr := postgres.WithInstance(db, &postgres.Config{})
	if postgresDriverErr != nil {
		return postgresDriverErr
	}

	migrationInstance, migrationInstanceErr := migrate.NewWithInstance("iofs", fileContent, databaseName, postgresDriver)
	if migrationInstanceErr != nil {
		return migrationInstanceErr
	}

	migrationErr := migrationInstance.Up()
	if migrationErr != nil {
		errMessage := migrationErr.Error()
		if errMessage == "no change" {
			return nil
		} else if scriptNr, scriptNrErr := IsErrDatabaseDirty(errMessage); scriptNr > 0 && scriptNrErr == nil {
			forcedMigrationErr := migrationInstance.Force(scriptNr)
			if forcedMigrationErr != nil {
				panic(forcedMigrationErr)
			}
		}
	}

	return nil
}

func IsErrDatabaseDirty(errorMessage string) (int, error) {
	regex, regexErr := regexp.Compile("Dirty database version ([0-9]+). Fix and force version.")
	if regexErr != nil {
		return 0, regexErr
	}

	dirtyMigrationNumberString := regex.Find([]byte(errorMessage))
	return cast.ToInt(dirtyMigrationNumberString), nil
}