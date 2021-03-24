package router

import (
	"helper/gin/controller/stock"

	"github.com/gin-gonic/gin"
)

func InitRouter(r *gin.Engine) {

	// v1 版本
	GroupV1 := r.Group("/v1")
	{
		GroupV1.Any("/stock/:symbol", stock.GetStock)
	}
}
