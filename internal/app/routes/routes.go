package routes

import (
	userHandler "terhandle/internal/features/user/handler"
	userRepository "terhandle/internal/features/user/repository"
	userService "terhandle/internal/features/user/service"

	"terhandle/internal/utils/jwt"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func UserRoute(e *echo.Echo, db *gorm.DB) {
	userRepository := userRepository.New(db)
	userService := userService.NewUserService(userRepository)
	userHandler := userHandler.NewUserHandler(userService)

	e.POST("/users", userHandler.Create)
	e.PUT("/users/:id", userHandler.CreateDetail, jwt.JWTMiddleware())
	e.PUT("/request/:id", userHandler.CreateTeknisiRole, jwt.JWTMiddleware())
	e.POST("/login", userHandler.Login)
	e.GET("/home", userHandler.SelectAll)


}
