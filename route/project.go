package route

import (
	"github.com/gin-gonic/gin"

	"github.com/Penetration-Platform-Go/Admin-Service/middleware"

	"github.com/Penetration-Platform-Go/Admin-Service/controller"
)

func projectServiceRoute(route *gin.RouterGroup) {

	route.GET("/", middleware.Auth(), controller.QueryAllProjects)
	route.DELETE("/", middleware.Auth(), controller.DeleteProject)

	route.PUT("/score", middleware.Auth(), controller.EvaluateProject)

	route.GET("/user", middleware.Auth(), controller.QueryProjectsByUser)
	route.GET("/id", middleware.Auth(), controller.QueryProjectByID)

}
