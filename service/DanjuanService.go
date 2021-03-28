package service

import (
	"fmt"
	"helper/api/danjuan"
)

func GetLowIndex() []danjuan.Items {
	items := danjuan.GetIndex()
	low := make([]danjuan.Items, 0) // len(items))
	for i := 0; i < len(items); i++ {
		if items[i].Roe > 0.15 && items[i].PePercentile < 0.30 {
			fmt.Println("name:", items[i].Name, " roe:", items[i].Roe, " PeOverHistory:", items[i].PeOverHistory)
			low = append(low, items[i])
		}
	}
	return low
}
