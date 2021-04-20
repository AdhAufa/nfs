package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// type CheckHandler interface {
// 	Health(ctx *gin.Context)
// }

// func NewCheckHandler() CheckHandler {

// }

func Health(ctx *gin.Context) {
	response := map[string]string{
		"message": "ok!",
	}
	ctx.JSON(http.StatusOK, response)
}
