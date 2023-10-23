package database

import (
	"terhandle/internal/app/model"
	"terhandle/internal/app/config"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// database connection
func InitDBMysql(cfg *config.AppConfig) *gorm.DB {

	// declare struct config & variable connectionString
	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		cfg.DBUSER, cfg.DBPASS, cfg.DBHOST, cfg.DBPORT, cfg.DBNAME)

	db, err := gorm.Open(mysql.Open(connectionString), &gorm.Config{})

	if err != nil {
		panic(err)
	}
	return db
}

func InitMigrationMysql(db *gorm.DB) {
	db.AutoMigrate(&model.Users{})
	db.AutoMigrate(&model.Achievement{})
	// db.AutoMigrate(&model.RequestTeknisi{})
}