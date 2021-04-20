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
	db             *gorm.DB               = config.SetupDatabaseConnection()
	userRepo       repo.UserRepository    = repo.NewUserRepo(db)
	productRepo    repo.ProductRepository = repo.NewProductRepo(db)
	authService    service.AuthService    = service.NewAuthService(userRepo)
	jwtService     service.JWTService     = service.NewJWTService()
	userService    service.UserService    = service.NewUserService(userRepo)
	productService service.ProductService = service.NewProductService(productRepo)
	authHandler    v1.AuthHandler         = v1.NewAuthHandler(authService, jwtService, userService)
	userHandler    v1.UserHandler         = v1.NewUserHandler(userService)
	productHandler v1.ProductHandler      = v1.NewProductHandler(productService)
)

func main() {
	defer config.CloseDatabaseConnection(db)
	server := gin.Default()

	authRoutes := server.Group("api/auth")
	{
		authRoutes.POST("/login", authHandler.Login)
		authRoutes.POST("/register", authHandler.Register)
	}

	userRoutes := server.Group("api/user")
	{
		userRoutes.GET("/profile", userHandler.Profile)
	}

	productRoutes := server.Group("api/product")
	{
		productRoutes.POST("/", productHandler.CreateProduct)
		productRoutes.GET("/:id", productHandler.FindOneProductByID)
	}

	checkRoutes := server.Group("api/check")
	{
		checkRoutes.GET("health", v1.Health)
	}

	server.Run()
}
