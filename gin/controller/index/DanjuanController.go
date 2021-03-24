package index

import (
	"helper/service"

	"github.com/gin-gonic/gin"
)

func GetLowIndex(c *gin.Context) {
	res := service.GetLowIndex()
	c.JSON(200, gin.H{
		"code":    0,
		"message": "sucess",
		"data":    res,
	})
}
