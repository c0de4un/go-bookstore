package bookstore

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = =

// - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -
// INTERFACES
// - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -

type Controller struct {
}

// - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -
// PUBLIC.METHODS
// - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -

func (controller *Controller) WriteResponse(status bool, payload interface{}, ctx *gin.Context) {
	ctx.JSON(http.StatusOK, payload)
}

// = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = = =
