package main

import (
	"github.com/Penetration-Platform-Go/Admin-Service/route"
	"github.com/gin-gonic/gin"
)

func main() {
	app := gin.Default()

	// 动态绑定路由
	route.AdminRoute(app)
	app.Run(":8003")
}
