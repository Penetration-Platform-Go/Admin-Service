package route

import (
	"github.com/Penetration-Platform-Go/Admin-Service/middleware"
	"github.com/gin-gonic/gin"
)

// AdminRoute 路由分组
func AdminRoute(app *gin.Engine) {
	app.Use(middleware.Cors())
	projectService := app.Group("/admin/project")
	projectServiceRoute(projectService)

	userService := app.Group("/admin/user")
	userServiceRoute(userService)

}
