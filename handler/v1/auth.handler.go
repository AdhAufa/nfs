package v1

import (
	"net/http"
	"strconv"

	"github.com/adhaufa/nfs/common/errors"
	"github.com/adhaufa/nfs/common/obj"
	"github.com/adhaufa/nfs/dto"
	"github.com/adhaufa/nfs/service"
	"github.com/gin-gonic/gin"
)

type AuthHandler interface {
	Login(ctx *gin.Context)
	Register(ctx *gin.Context)
}

type authHandler struct {
	authService service.AuthService
	jwtService  service.JWTService
	userService service.UserService
}

func NewAuthHandler(
	authService service.AuthService,
	jwtService service.JWTService,
	userService service.UserService,
) AuthHandler {
	return &authHandler{
		authService: authService,
		jwtService:  jwtService,
		userService: userService,
	}
}

func (c *authHandler) Login(ctx *gin.Context) {
	var loginRequest dto.LoginRequest
	err := ctx.ShouldBind(&loginRequest)

	if err != nil {
		response := errors.BuildErrorResponse("Failed to process request", err.Error(), obj.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	err = c.authService.VerifyCredential(loginRequest.Email, loginRequest.Password)
	if err != nil {
		response := errors.BuildErrorResponse("Failed to login", err.Error(), obj.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
		return
	}

	user, _ := c.userService.FindUserByEmail(loginRequest.Email)

	token := c.jwtService.GenerateToken(strconv.FormatInt(user.ID, 64))
	user.Token = token
	response := errors.BuildResponse(true, "OK!", user)
	ctx.JSON(http.StatusOK, response)

}

func (c *authHandler) Register(ctx *gin.Context) {
	var registerRequest dto.RegisterRequest

	err := ctx.ShouldBind(&registerRequest)
	if err != nil {
		response := errors.BuildErrorResponse("Failed to process request", err.Error(), obj.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	user, err := c.userService.CreateUser(registerRequest)
	if err != nil {
		response := errors.BuildErrorResponse("Something gone wrong", err.Error(), obj.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	token := c.jwtService.GenerateToken(strconv.FormatInt(user.ID, 64))
	user.Token = token
	response := errors.BuildResponse(true, "OK!", user)
	ctx.JSON(http.StatusOK, response)

}
