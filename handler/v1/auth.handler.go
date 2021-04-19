package v1

import (
	"github.com/adhaufa/nfs/service"
	"github.com/gin-gonic/gin"
)

type AuthHandler interface {
	Login()
}

type authHandler struct {
	authService service.AuthService
	// jwtService  service.JWTService
}

func NewAuthHandler(authService service.AuthService) *authHandler {
	return &authHandler{
		authService: authService,
	}
}

func (c *authHandler) Login(ctx *gin.Context) {

}
