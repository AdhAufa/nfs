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
	authHandler v1.authHandler      = v1.NewAuthHandler(authService)
)

func main() {
	defer config.CloseDatabaseConnection(db)
	server := gin.Default()

	authRoutes := server.Group("api/auth")
	{
		authRoutes.POST("login")
		authRoutes.POST("register")
	}

	server.Run()
}
