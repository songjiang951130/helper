package mapper

import (
	"helper/dao/model"
	"helper/dao/mysql"
)

func GetSubscribeList() *[]model.StockSubscribe {
	db := mysql.Connect()
	stockSubscribe := new(model.StockSubscribe)
	res := make([]model.StockSubscribe, 0, 50)
	db.Model(&stockSubscribe).Find(&res)
	return &res
}
