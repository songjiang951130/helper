package main

import (
	"helper/dao/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "dream:dream1234@tcp(localhost:3306)/dream_go?charset=utf8mb4"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	//先删除表再创建表
	db.Migrator().DropTable(&model.StockQ4{})
	// 迁移 schema
	db.AutoMigrate(&model.StockQ4{})

	// // Create
	// db.Create(&Product{Code: "D42", Price: 100})

	// // Read
	// var product Product
	// db.First(&product, 1)                 // 根据整形主键查找
	// db.First(&product, "code = ?", "D42") // 查找 code 字段值为 D42 的记录

	// // Update - 将 product 的 price 更新为 200
	// db.Model(&product).Update("Price", 200)
	// // Update - 更新多个字段
	// db.Model(&product).Updates(Product{Price: 200, Code: "F42"}) // 仅更新非零值字段
	// db.Model(&product).Updates(map[string]interface{}{"Price": 200, "Code": "F42"})

	// // Delete - 删除 product
	// db.Delete(&product, 1)
}
