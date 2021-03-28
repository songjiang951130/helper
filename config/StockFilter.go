package config

var StockMap = map[string]string{
	".NDX":     "纳斯达克100指数",
	".INX":     "标普500指数",
	"SZ399989": "中证医疗",
	"SH512120": "医药50ETF",
	"SH513050": "中概互联网ETF",
	"SH515790": "光伏ETF",
	"SH000300": "沪深300",
}

func GetStockFilter() map[string]string {
	return StockMap
}
