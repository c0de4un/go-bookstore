package bookstore

// = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = =

// - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -
// IMPORTS
// - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"

	database "github.com/c0de4un/go-bookstore/core/database"
)

// - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -
// PUBLIC.FUNCTIONS
// - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -

func Run() {
	fmt.Println("Loading database")
	db, err := database.LoadDatabase("")
	if err != nil {
		log.Fatalf("\n bookstore::Run: %s \n", err)
	}

	fmt.Println("Connecting to Database")
	err = db.Connect()
	if err != nil {
		log.Fatalf("\n bookstore::Run: %s \n", err)
	}
	defer db.Disconnect()

	fmt.Println("Initializing database")
	err = initializeDB(db)
	if err != nil {
		log.Fatalf("\n bookstore::Run: %s \n", err)
	}

	fmt.Println("Initializing Router")
	router := gin.Default()

	initializeRoutes(router)

	router.Run()
}

// = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = =
