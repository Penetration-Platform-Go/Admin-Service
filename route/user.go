package route

import (
	"github.com/gin-gonic/gin"

	"github.com/Penetration-Platform-Go/Admin-Service/middleware"

	"github.com/Penetration-Platform-Go/Admin-Service/controller"
)

func userServiceRoute(route *gin.RouterGroup) {

	route.GET("/", middleware.Auth(), controller.QueryAllUsers)

	route.GET("/username", middleware.Auth(), controller.QueryUserByUsername)
}
