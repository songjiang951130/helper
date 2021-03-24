package main

import (
	"helper/gin/router"

	"github.com/gin-gonic/gin"
)

func main() {
	// gin.SetMode(gin.ReleaseMode)
	// 默认启动方式，包含 Logger、Recovery 中间件
	engine := gin.Default()
	router.InitRouter(engine)
	//visit http://localhost:8082/ping
	engine.Run(":8082")
}
