package routes

import (
	"terhandle/internal/features/user/handler"
	"terhandle/internal/features/user/repository"
	"terhandle/internal/features/user/service"
	"terhandle/internal/utils/jwt"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func UserRoute (e *echo.Echo,db *gorm.DB){
	userRepository := repository.New(db)
	userService := service.NewUserService(userRepository)
	userHandler := handler.NewUserHandler(userService)

	e.POST("/users", userHandler.Create)
	e.PUT("/users/:id", userHandler.CreateDetail,jwt.JWTMiddleware())
	e.PUT("/request/:id", userHandler.CreateTeknisiRole,jwt.JWTMiddleware())
	// e.GET("/users", userController.SelectAll,middleware.JWTMiddleware())
	e.POST("/login",userHandler.Login)
}