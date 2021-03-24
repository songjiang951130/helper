package service

import (
	"fmt"
	"helper/api/danjuan"
)

func GetLowIndex() []danjuan.Items {
	items := danjuan.GetIndex()
	typeStr := fmt.Sprintf("%T", items)
	fmt.Println(typeStr)
	low := make([]danjuan.Items, len(items))
	for i := 0; i < len(items); i++ {
		fmt.Println(items[i])
		fmt.Println("roe:", items[i].Roe, " PeOverHistory:", items[i].PeOverHistory)
		if items[i].Roe > 0.15 && items[i].PeOverHistory < 0.3 {
			size := len(low)
			fmt.Println("size:", size)
			var index int
			if size == 0 {
				index = 0
			} else {
				index = size - 1
			}
			low[index] = items[i]
		}
	}
	return low
}
