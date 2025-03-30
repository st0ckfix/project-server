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
	deviceRepo := repository.NewDeviceRepository(config.DB)
	scheduleRepo := repository.NewScheduleRepository(config.DB)

	// Initialize usecase
	userUsecase := usecase.NewUserUsecase(userRepo, logRepo)
	profileUsecase := usecase.NewProfileUsecase(profileRepo)
	deviceUsecase := usecase.NewDeviceUsecase(deviceRepo)
	scheduleUsecase := usecase.NewScheduleUsecase(scheduleRepo)

	// Initialize handler
	userHandler := handler.NewUserHandler(userUsecase)
	profileHandler := handler.NewProfileHandler(profileUsecase)
	deviceHandler := handler.NewDeviceHandler(deviceUsecase)
	scheduleHandler := handler.NewScheduleHandler(scheduleUsecase)

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

	device := apiV1.Group("/device")
	{
		// Device routes
		device.POST("/add", deviceHandler.AddDevice)
		device.GET("/get", deviceHandler.GetDevices)
		device.PUT("/update", deviceHandler.UpdateDevice)
		device.DELETE("/remove", deviceHandler.RemoveDevice)
	}

	schedule := apiV1.Group("/schedule")
	{
		// Schedule routes
		schedule.POST("/add", scheduleHandler.AddSchedule)
		schedule.GET("/get", scheduleHandler.GetSchedules)
		schedule.PUT("/update", scheduleHandler.UpdateSchedule)
		schedule.DELETE("/remove", scheduleHandler.RemoveSchedule)
	}

	// Start server
	_ = r.Run("0.0.0.0:8085") // Chạy trên tất cả các địa chỉ mạng
}
