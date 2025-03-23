package main

import (
	"fastbuy/config"
	"fastbuy/internal/handler"
	"fastbuy/internal/middleware"
	"fastbuy/internal/repository"
	"fastbuy/internal/usecase"
	"github.com/gin-gonic/gin"
)

func main() {
	// Load environment variables
	config.LoadEnv()

	// Connect to PostgreSQL and MongoDB
	config.ConnectPostgres()
	config.ConnectMongoDB()

	// Initialize repositories
	userRepo := repository.NewUserRepository(config.DB)
	profileRepo := repository.NewProfileRepository(config.DB)
	logRepo := repository.NewMongoRepository(config.MongoDB)

	// Initialize usecase
	userUsecase := usecase.NewUserUsecase(userRepo, logRepo)
	profileUsecase := usecase.NewProfileUsecase(profileRepo)

	// Initialize handler
	userHandler := handler.NewUserHandler(userUsecase)
	profileHandler := handler.NewProfileHandler(profileUsecase)

	// Setup Gin router
	r := gin.Default()

	// Group API v1
	apiV1 := r.Group("/api/v1")

	// Group Authentication
	auth := apiV1.Group("/auth")
	{
		// User routes
		auth.POST("/register", userHandler.Register)
		auth.POST("/login", userHandler.Login)
		auth.POST("/refresh_token", userHandler.RefreshToken)
		auth.POST("/logout", userHandler.Logout)
	}

	profile := apiV1.Group("/profile")
	{
		// Profile routes
		profile.GET("/details", middleware.AuthMiddleware(), profileHandler.GetProfile)
		profile.PUT("/update", middleware.AuthMiddleware(), profileHandler.UpdateProfile)
	}

	// Start server
	_ = r.Run("0.0.0.0:8085") // Chạy trên tất cả các địa chỉ mạng
}
