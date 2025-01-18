package migrator

import (
	"database/sql"

	"github.com/pressly/goose/v3"
)

const (
	postgresDriver    = "postgres"
	migrationFilePath = "/schema/migrations"
)

// RunMigrations performs the migrations for the migration files
func RunMigrations(db *sql.DB) error {
	err := goose.SetDialect(postgresDriver)
	if err != nil {
		return err
	}

	goose.SetTableName("db_version_account_transactions")

	err = goose.Up(db, migrationFilePath)
	if err != nil {
		return err
	}

	return nil
}
