package v1

import (
	"github.com/adhaufa/nfs/service"
	"github.com/gin-gonic/gin"
)

type UserHandler interface {
	Profile(ctx *gin.Context)
}

type userHandler struct {
	userService service.UserService
}

func NewUserHandler(userService service.UserService) UserHandler {
	return &userHandler{
		userService: userService,
	}
}

func (c *userHandler) Profile(ctx *gin.Context) {
	user, err := c.userService.FindUserByEmail()
}
