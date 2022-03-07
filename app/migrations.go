package bookstore

// = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = =

// - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -
// IMPORTS
// - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -

import (
	migrate "github.com/c0de4un/go-bookstore/app/migrations"
	database "github.com/c0de4un/go-bookstore/core/database"
)

// - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -
// PUBLIC.FUNCTIONS
// - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -

func initializeDB(db *database.Database) error {
	err := db.RunMigration(&migrate.AddAccounts{})
	if err != nil {
		return err
	}

	return nil
}

// = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = =
