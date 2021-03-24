package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type StockQ4 struct {
	ID         uint32  `gorm:"primaryKey "`
	Symbol     string  `gorm:"type:char(16) CHARACTER SET utf8 COLLATE utf8_bin NOT NULL DEFAULT '' COMMENT '代码';index:symbol_idx"`
	Name       string  `gorm:"type:varchar(16) CHARACTER SET utf8 COLLATE utf8_bin NOT NULL DEFAULT '' COMMENT '名字'"`
	ReportName string  `gorm:"type:varchar(16) CHARACTER SET utf8 COLLATE utf8_bin NOT NULL DEFAULT '' COMMENT '报告期名字'"`
	ReportTime uint    `gorm:"type:int unsigned NOT NULL DEFAULT '0' COMMENT '报告期'"`
	Roe        float64 `gorm:"type:double(11,4) NOT NULL DEFAULT '0.0000' COMMENT '净资产收益率'"`
	NetSell    float64 `gorm:"type:double(11,4) NOT NULL DEFAULT '0.0000' COMMENT '净利率'"`
	GrossSell  float64 `gorm:"type:double(11,4) NOT NULL DEFAULT '0.0000' COMMENT '毛利率'"`
	AssertRate float64 `gorm:"type:double(11,4) NOT NULL DEFAULT '0.0000' COMMENT '资产负债率'"`
	Type       string  `gorm:"type:char(2) CHARACTER SET utf8 COLLATE utf8_bin NOT NULL DEFAULT ''"`
}

func main() {
	dsn := "dream:dream1234@tcp(localhost:3306)/dream_go?charset=utf8mb4"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	//先删除表再创建表
	db.Migrator().DropTable(&StockQ4{})
	// 迁移 schema
	db.AutoMigrate(&StockQ4{})

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
