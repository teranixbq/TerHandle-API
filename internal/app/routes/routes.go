package routes

import (
	// User
	userHandler "terhandle/internal/features/user/handler"
	userRepository "terhandle/internal/features/user/repository"
	userService "terhandle/internal/features/user/service"

	// Request Teknisi
	requestHandler "terhandle/internal/features/request-teknisi/handler"
	requestRepository "terhandle/internal/features/request-teknisi/repository"
	requestService "terhandle/internal/features/request-teknisi/service"

	// Feedback
	feedbackHandler "terhandle/internal/features/feedback/handler"
	feedbackRepository "terhandle/internal/features/feedback/repository"
	feedbackService "terhandle/internal/features/feedback/service"

	// Admin
	adminHandler "terhandle/internal/features/admin/handler"
	adminRepository "terhandle/internal/features/admin/repository"
	adminService "terhandle/internal/features/admin/service"

	"terhandle/internal/utils/jwt"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func InitRoute(e *echo.Echo, db *gorm.DB) {
	// User
	userRepository := userRepository.New(db)
	userService := userService.NewUserService(userRepository)
	userHandler := userHandler.NewUserHandler(userService)

	// Request Teknisi
	requestRepository := requestRepository.New(db)
	requestService := requestService.NewUserService(requestRepository)
	requestHandler := requestHandler.NewUserHandler(requestService)

	// Feedback
	feedbackRepository := feedbackRepository.New(db)
	feedbackService := feedbackService.NewUserService(feedbackRepository)
	feedbackHandler := feedbackHandler.NewUserHandler(feedbackService)

	// Admin
	adminRepository := adminRepository.New(db)
	adminService := adminService.NewUserService(adminRepository)
	adminHandler := adminHandler.NewUserHandler(adminService)

	e.POST("/register", userHandler.Create)
	e.POST("/login", userHandler.Login)
	e.GET("/", userHandler.GetAll)
	e.GET("/:id", userHandler.GetByIdWithFeedback)
	e.PUT("/teknisi/register/:id", userHandler.CreateTeknisiRole, jwt.JWTMiddleware())

	user := e.Group("/profile", jwt.JWTMiddleware())
	user.GET("",userHandler.GetById)
	user.PUT("/:id", userHandler.CreateDetail)
	user.DELETE("/:id", userHandler.DeleteById)

	request := e.Group("/request", jwt.JWTMiddleware())
	request.POST("", requestHandler.Create)
	request.GET("/history/:id", requestHandler.GetAllHistoryRequest)
	request.GET("/history/:id/:id_request", requestHandler.GetHistoryRequestById)
	request.PUT("/:id/:id_request", requestHandler.UpdateStatusRequest)

	feedback := e.Group("/feedback", jwt.JWTMiddleware())
	feedback.POST("/:id", feedbackHandler.Create)
	feedback.PUT("/:id/:id_feedback", feedbackHandler.UpdateDataFeedback)

	admin := e.Group("/", jwt.JWTMiddleware())
	admin.POST("profesi",adminHandler.Create)
	admin.GET("profesi",adminHandler.GetAllProfesi)
	admin.PUT("profesi/:id_profesi",adminHandler.UpdateDataProfesi)
	admin.DELETE("profesi/:id_profesi",adminHandler.DeleteDataProfesi)
	admin.POST("transport",adminHandler.CreateBiaya)

}
