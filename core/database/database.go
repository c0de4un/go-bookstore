package bookstore

// = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = =

// - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -
// IMPORTS
// - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -

// - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -
// STRUCTS
// - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -

type Database struct {
	name     string
	host     string
	port     int
	login    string
	password string
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

	return db, nil
}

// = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = =
