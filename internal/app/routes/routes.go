package routes

import (
	userHandler "terhandle/internal/features/user/handler"
	userRepository "terhandle/internal/features/user/repository"
	userService "terhandle/internal/features/user/service"

	requestHandler "terhandle/internal/features/request-teknisi/handler"
	requestRepository "terhandle/internal/features/request-teknisi/repository"
	requestService "terhandle/internal/features/request-teknisi/service"

	feedbackHandler "terhandle/internal/features/feedback/handler"
	feedbackRepository "terhandle/internal/features/feedback/repository"
	feedbackService "terhandle/internal/features/feedback/service"

	"terhandle/internal/utils/jwt"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func UserRoute(e *echo.Echo, db *gorm.DB) {
	userRepository := userRepository.New(db)
	userService := userService.NewUserService(userRepository)
	userHandler := userHandler.NewUserHandler(userService)

	requestRepository := requestRepository.New(db)
	requestService := requestService.NewUserService(requestRepository)
	requestHandler := requestHandler.NewUserHandler(requestService)

	feedbackRepository := feedbackRepository.New(db)
	feedbackService := feedbackService.NewUserService(feedbackRepository)
	feedbackHandler := feedbackHandler.NewUserHandler(feedbackService)

	e.POST("/register", userHandler.Create)
	e.POST("/login", userHandler.Login)
	e.GET("/", userHandler.SelectAll)
	e.GET("/:id",userHandler.SelectById)
	e.PUT("/users/:id", userHandler.CreateDetail, jwt.JWTMiddleware())
	e.DELETE("/users/:id", userHandler.DeleteById, jwt.JWTMiddleware())
	e.PUT("/teknisi/register/:id", userHandler.CreateTeknisiRole, jwt.JWTMiddleware())

	e.POST("/request", requestHandler.Create, jwt.JWTMiddleware())
	e.GET("/request/history/:id", requestHandler.GetAllHistoryRequest, jwt.JWTMiddleware())
	e.GET("/request/:id/:id_request", requestHandler.GetHistoryRequestById, jwt.JWTMiddleware())
	e.PUT("/request/:id/:id_request", requestHandler.UpdateStatusRequest, jwt.JWTMiddleware())
	//e.PUT("/request/:id", requestHandler.Create, jwt.JWTMiddleware())

	e.POST("/feedback/:id", feedbackHandler.Create, jwt.JWTMiddleware())
	e.PUT("/feedback/:id/:id_feedback", feedbackHandler.UpdateDataFeedback, jwt.JWTMiddleware())
	

}
