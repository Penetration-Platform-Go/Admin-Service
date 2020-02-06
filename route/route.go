package route

import "github.com/gin-gonic/gin"

// AdminRoute 路由分组
func AdminRoute(app *gin.Engine) {
	projectService := app.Group("/admin/project")
	projectServiceRoute(projectService)

	userService := app.Group("/admin/user")
	userServiceRoute(userService)

}
