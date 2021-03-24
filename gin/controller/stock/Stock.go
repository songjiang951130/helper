package stock

import "github.com/gin-gonic/gin"

func GetStock(c *gin.Context) {
	// 获取 Get 参数
	symbol := c.Param("symbol")

	c.JSON(200, gin.H{
		"symbol": symbol,
		"name":   "sss",
	})
}
