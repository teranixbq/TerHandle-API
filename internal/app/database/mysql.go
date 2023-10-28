package database

import (
	"fmt"
	"terhandle/internal/app/config"
	"terhandle/internal/app/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDBMysql(cfg *config.AppConfig) *gorm.DB {

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
	db.AutoMigrate(&model.RequestTeknisi{})
	db.AutoMigrate(&model.Foto{})
	db.AutoMigrate(&model.Feedback{})

}