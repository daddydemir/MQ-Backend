package config

import (
	"github.com/daddydemir/MQ/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func GetConnect() {
	dsn := "root:pass@tcp(127.0.0.1:3306)/test_db?parseTime=True"
	DB, _ = gorm.Open(mysql.Open(dsn), &gorm.Config{})
}

func CreateTables() {
	// db tables initial functions...
	_ = DB.AutoMigrate(model.Person{}, model.History{}, model.Question{}, model.Answer{})
}
