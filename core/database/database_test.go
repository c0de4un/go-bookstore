package bookstore_test

// = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = =

// - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -
// IMPORTS
// - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -

import (
	"testing"

	database "github.com/c0de4un/go-bookstore/core/database"
)

// - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -
// UNITS
// - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -

func TestLoadDatabase(t *testing.T) {
	db, err := database.LoadDatabase("../../config/database.example.xml")
	if err != nil {
		t.Errorf("Database::LoadDatabase: error: '%s'", err)
	}

	if db.GetName() != "bookstore" {
		t.Error("Database::LoadDatabase: deserialization failed")
	}
}

// = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = =
