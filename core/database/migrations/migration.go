package bookstore

// = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = =

// - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -
// IMPORTS
// - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -

import (
	"database/sql"
)

// - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -
// INTERFACES
// - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -

type IMigration interface {
	GetID() int
	GetName() string
	Up(db *sql.DB) error
	Down(db *sql.DB) error
}

// - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -
// STRUCTS
// - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -

type Migration struct {
}

// = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = =
