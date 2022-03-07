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
	db, err := database.LoadDatabase("")
	if err != nil {
		log.Fatalf("\n bookstore::Run: %s \n", err)
	}

	err = db.Connect()
	if err != nil {
		log.Fatalf("\n bookstore::Run: %s \n", err)
	}
	defer db.Disconnect()

	router := gin.Default()

	initializeRoutes(router)

	fmt.Println("Database connection test complete")
	// router.Run()
}

// = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = =
