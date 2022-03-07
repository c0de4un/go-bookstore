package bookstore

// = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = =

// - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -
// IMPORTS
// - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -

import (
	"database/sql"

	migrations "github.com/c0de4un/go-bookstore/core/database/migrations"
)

// - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -
// STRUCTS
// - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -

type AddAccounts struct {
	migrations.Migration
}

// - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -
// PUBLIC.METHODS
// - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -

func NewAddAccounts() *AddAccounts {
	return &AddAccounts{}
}

// - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -
// IMigration
// - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -

func (instance *AddAccounts) GetID() int {
	return 1
}

func (instance *AddAccounts) GetName() string {
	return "accounts"
}

func (instance *AddAccounts) Up(db *sql.DB) error {
	_, err := db.Exec("CREATE TABLE `accounts` (`id` BIGINT AUTO_INCREMENT PRIMARY KEY, `login` VARCHAR(32) NOT NULL, `email` VARCHAR(128) NOT NULL, `name` VARCHAR(64) NULL, `surname` VARCHAR(64), `password` VARCHAR(128) NOT NULL, `created_at` DATETIME NOT NULL, `updated_at` DATETIME NOT NULL)")
	if err != nil {
		return err
	}

	return nil
}

func (instance *AddAccounts) Down(db *sql.DB) error {
	_, err := db.Exec("DROP TABLE `accounts`")
	if err != nil {
		return err
	}

	return nil
}

// = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = =
