package main

import "helper/api/xueqiu"

func main() {
	symbol := "sz000001"
	xueqiu.GetTopHolder(symbol)
}
