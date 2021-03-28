package service

import (
	"fmt"
	"helper/api/danjuan"
	"helper/config"
	"helper/dao/mapper"
	"helper/dao/model"
	"sync"
	"time"

	"github.com/shopspring/decimal"
)

var lowPeList []LowPe
var mu sync.Mutex

type User struct {
	Name      string
	Email     string
	Git       string
	Dan       []danjuan.Items
	LowPeList []LowPe
}

type LowPe struct {
	Name      string
	Symbol    string
	Pe        string
	StartTime string
	EndTime   string
	Rate      float64
	RateStr   string
}

func GetIndexInfo() User {
	dan := GetLowIndex()
	lowPe := getSingleLowPeList()
	u := User{"宋江", "songjiangcq@qq.com", "https://gitee.com/songjiang_951130/helper.git", dan, lowPe}
	return u
}

func getSingleLowPeList() []LowPe {
	if lowPeList == nil {
		mu.Lock()
		defer mu.Unlock()
		if lowPeList == nil {
			fmt.Println("创建单例")
			lowPeList = GetLowPe()
		}
	}
	return lowPeList
}
func GetLowPe() []LowPe {
	sub := mapper.GetSubscribeList()
	subList := *sub
	res := make([]LowPe, 0, 20)
	stockFilter := config.GetStockFilter()
	for i := 0; i < len(subList); i++ {
		stock := subList[i]
		_, ok := stockFilter[stock.Symbol]
		if ok {
			continue
		}
		p := mapper.GetDailyList(stock.Symbol)
		low := cacl(p)
		if low.Rate < 25 {
			low.Name = stock.Name
			low.Symbol = stock.Symbol
			res = append(res, low)
		}
	}
	return res
}

func cacl(data *[]model.StockDaily) LowPe {
	dataList := *data
	lowPe := LowPe{Rate: 100.0}
	if len(dataList) < 2 {
		lowPe.Rate = 100.0
		return lowPe
	}
	start := dataList[0].Timestamp
	fmt.Println("start:", start)
	startStr := time.Unix(start/1000, 0).Format("2006-01-02")
	lowPe.StartTime = startStr
	end := dataList[len(dataList)-1].Timestamp
	fmt.Println("end:", end)
	endStr := time.Unix(end/1000, 0).Format("2006-01-02")
	lowPe.EndTime = endStr

	lastPe := dataList[len(dataList)-1].Pe
	var rank int64 = 0
	for i := 0; i < len(dataList); i++ {
		t := dataList[i]
		if lastPe > t.Pe {
			rank++
		}
	}
	rate, _ := decimal.NewFromInt(100).Mul(decimal.NewFromInt(rank)).Div(decimal.NewFromInt(int64(len(dataList)))).Float64()
	lowPe.Rate = rate
	lowPe.RateStr = fmt.Sprintf("%0.2f", rate)
	lowPe.Pe = fmt.Sprintf("%0.2f", lastPe)
	fmt.Println("lowPe:", lowPe)
	return lowPe
}
