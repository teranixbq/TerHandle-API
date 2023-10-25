package main

import (
	"fmt"
	"terhandle/internal/app/config"
	"terhandle/internal/app/database"
	"terhandle/internal/app/routes"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	cfg := config.InitConfig()
	db := database.InitDBMysql(cfg)
	database.InitMigrationMysql(db)

	routes.UserRoute(e, db)
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: `[${time_rfc3339}] ${status} ${method} ${host}${path} ${latency_human}` + "\n",
	}))

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", cfg.SERVERPORT)))
}
