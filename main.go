package main

import (
	"github.com/adhaufa/nfs/config"
	v1 "github.com/adhaufa/nfs/handler/v1"
	"github.com/adhaufa/nfs/repo"
	"github.com/adhaufa/nfs/service"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var (
	db          *gorm.DB            = config.SetupDatabaseConnection()
	userRepo    repo.UserRepository = repo.NewUserRepo(db)
	authService service.AuthService = service.NewAuthService(userRepo)
	jwtService  service.JWTService  = service.NewJWTService()
	userService service.UserService = service.NewUserService(userRepo)
	authHandler v1.AuthHandler      = v1.NewAuthHandler(authService, jwtService, userService)
)

func main() {
	defer config.CloseDatabaseConnection(db)
	server := gin.Default()

	authRoutes := server.Group("api/auth")
	{
		authRoutes.POST("login", authHandler.Login)
		authRoutes.POST("register", authHandler.Register)
	}

	checkRoutes := server.Group("api/check")
	{
		checkRoutes.GET("health", v1.Health)
	}

	server.Run()
}
