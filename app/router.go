package bookstore

// = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = =

// - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -
// IMPORTS
// - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -

import (
	"github.com/gin-gonic/gin"

	auth "github.com/c0de4un/go-bookstore/app/http/controllers/auth"
)

// - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -
// PUBLIC.FUNCTIONS
// - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -

func initializeRoutes(router *gin.Engine) {
	auth := auth.New()

	router.GET("/api/v1/auth/register", auth.RegisterHandler)
}

// = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = =
