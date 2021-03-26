package router

import (
	"helper/gin/controller/index"
	"helper/gin/controller/stock"

	"github.com/gin-gonic/gin"
)

const server_context string = "/go"

func InitRouter(r *gin.Engine) {
	loadTemplate(r)
	r.GET(server_context, index.GetIndex)

	// v1 版本
	GroupV1 := r.Group(server_context + "/v1")
	{
		//参数path和 handler
		GroupV1.Any("/stock/:symbol", stock.GetStock)
		GroupV1.GET("/index", index.GetLowIndex)
	}
}

func loadTemplate(r *gin.Engine) {
	r.LoadHTMLGlob("./resource/html/*.html")
}
