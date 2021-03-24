package router

import (
	"helper/gin/controller/index"
	"helper/gin/controller/stock"

	"github.com/gin-gonic/gin"
)

func InitRouter(r *gin.Engine) {

	// v1 版本
	GroupV1 := r.Group("/go/v1")
	{
		//参数path和 handler
		GroupV1.Any("/stock/:symbol", stock.GetStock)
		GroupV1.Any("/index", index.GetLowIndex)
	}
}
