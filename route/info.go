package route

import "github.com/gin-gonic/gin"

import "github.com/Penetration-Platform-Go/Admin-Service/middleware"

import "github.com/Penetration-Platform-Go/Admin-Service/controller"

func mainServiceRoute(route *gin.RouterGroup) {
	route.GET("/", middleware.Auth(), controller.GetInfo)
}
