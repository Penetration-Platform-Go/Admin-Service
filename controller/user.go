package controller

import (
	"github.com/Penetration-Platform-Go/Admin-Service/model"
	"github.com/gin-gonic/gin"
)

// QueryAllUsers handle
func QueryAllUsers(ctx *gin.Context) {
	result, err := model.QueryAllUsers()
	if err != nil {
		ctx.Status(400)
	} else {
		ctx.JSON(200, result)
	}
}

// QueryUserByUsername handle
func QueryUserByUsername(ctx *gin.Context) {
	result, err := model.QueryUserByUsername(ctx.Query("username"))
	if err != nil {
		ctx.Status(400)
	} else {
		ctx.JSON(200, result)
	}
}
