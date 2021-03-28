package mapper

import (
	"helper/dao/model"
	"helper/dao/mysql"
)

func GetDailyList(symbol string) *[]model.StockDaily {
	db := mysql.Connect()
	selectValue := new(model.StockDaily)
	res := make([]model.StockDaily, 0, 300)
	db.Model(&selectValue).Where("symbol = ?", symbol).Order("timestamp").Limit(500).Find(&res)
	return &res
}
