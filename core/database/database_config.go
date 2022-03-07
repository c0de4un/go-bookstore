package bookstore

// = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = =

// - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -
// IMPORTS
// - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -

import (
	"encoding/xml"
	"io/ioutil"
	"os"
)

// - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -
// CONSTANTS
// - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -

const db_config_path string = "./config/database.xml"

// - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -
// STRUCTS
// - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -

type DatabaseXML struct {
	XMLName  xml.Name `xml:"DatabaseConfig"`
	Name     string   `xml:"name,attr"`
	Host     string   `xml:"host,attr"`
	Port     int      `xml:"port,attr"`
	Login    string   `xml:"login,attr"`
	Password string   `xml:"password,attr"`
}

// - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -
// PUBLIC.METHODS
// - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -

func loadDatabaseXML(path string) (*DatabaseXML, error) {
	var configPath string
	if len(path) > 0 {
		configPath = path
	} else {
		configPath = db_config_path
	}

	file, err := os.Open(configPath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}

	dbXML := &DatabaseXML{}
	err = xml.Unmarshal(bytes, dbXML)
	if err != nil {
		return dbXML, err
	}

	return dbXML, nil
}

// = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = =
