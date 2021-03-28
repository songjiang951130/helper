package mysql

import (
	"helper/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Connect() *gorm.DB {
	dsn := "dream:dream1234@tcp(localhost:3306)/dream?charset=utf8mb4"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	return db
}

func getConfig() {
	config.GetDataSourceConfig()
}
