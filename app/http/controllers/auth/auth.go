package bookstore

// = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = =

// - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -
// IMPORTS
// - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -

import (
	"net/http"

	"github.com/gin-gonic/gin"

	base_controller "github.com/c0de4un/go-bookstore/core/http/controllers"
)

// - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -
// STRUCTS
// - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -

type AuthController struct {
	base_controller.Controller
}

// - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -
// PUBLIC.FUNCTIONS
// - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -

func New() *AuthController {
	return &AuthController{}
}

func (controller *AuthController) RegisterHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"status": true, "message": "authentication required"})

	// @TODO: AuthController::RegisterHandler()
}

// = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = =
