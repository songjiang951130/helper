package main

import (
	"fmt"
	"helper/gin/router"

	"github.com/gin-gonic/gin"
)

func main() {
	// gin.SetMode(gin.ReleaseMode)
	// 默认启动方式，包含 Logger、Recovery 中间件
	engine := gin.Default()
	router.InitRouter(engine)
	//visit http://localhost:2021/go
	fmt.Println("http://localhost:2021/go")
	engine.Run(":2021")
}
