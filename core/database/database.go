package bookstore

// = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = =

// - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -
// IMPORTS
// - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -

import (
	"database/sql"
	"errors"
	"fmt"

	migrations "github.com/c0de4un/go-bookstore/core/database/migrations"
	_ "github.com/go-sql-driver/mysql"
)

// - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -
// STRUCTS
// - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -

type Database struct {
	name     string
	host     string
	port     int
	login    string
	password string
	db       *sql.DB
}

// - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -
// GETTERS & SETTERS
// - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -

func (db *Database) GetName() string {
	return db.name
}

func (db *Database) GetHost() string {
	return db.host
}

func (db *Database) GetPort() int {
	return db.port
}

func (db *Database) GetLogin() string {
	return db.login
}

// - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -
// PUBLIC.METHODS
// - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -

func LoadDatabase(path string) (*Database, error) {
	dbXML, err := loadDatabaseXML(path)
	if err != nil {
		return nil, err
	}

	db := &Database{}
	db.name = dbXML.Name
	db.host = dbXML.Host
	db.port = dbXML.Port
	db.login = dbXML.Login
	db.password = dbXML.Password

	if len(db.login) < 1 {
		return db, errors.New("Database::LoadDatabase: invalid config. Check xml")
	}

	fmt.Printf("\n Database:\n name='%s'\n login='%s' \n", db.name, db.login)

	return db, nil
}

func (instance *Database) RunMigration(migration migrations.IMigration) error {
	fmt.Printf("\n Database::RunMigration: '%d' \n", migration.GetID())

	return migration.Up(instance.db)
}

func (instance *Database) Connect() error {
	if len(instance.login) < 1 {
		return errors.New("Database::Connect: invalid config. Check xml")
	}

	src := fmt.Sprintf("%s:%s@/%s", instance.login, instance.password, instance.name)
	db, err := sql.Open("mysql", src)
	if err != nil {
		return err
	}
	instance.db = db

	return nil
}

func (instance *Database) Disconnect() error {
	return instance.db.Close()
}

// = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = =
