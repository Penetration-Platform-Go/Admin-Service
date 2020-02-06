package controller

import (
	"fmt"

	"github.com/Penetration-Platform-Go/Admin-Service/model"
	"github.com/gin-gonic/gin"
)

// QueryAllProjects handle
func QueryAllProjects(ctx *gin.Context) {
	project, err := model.QueryAllProjects()
	if err != nil {
		ctx.Status(400)
	} else {
		ctx.JSON(200, project)
	}
}

// DeleteProject handle
func DeleteProject(ctx *gin.Context) {
	flag, result := model.DeleteProjectByID(ctx.Query("id"))
	if flag {
		ctx.Status(200)
	} else {
		ctx.String(400, result)
	}
}

// EvaluateProject handle
func EvaluateProject(ctx *gin.Context) {
	var project model.Project
	err := ctx.BindJSON(&project)
	if err != nil {
		fmt.Println(err)
		ctx.Status(406)
		return
	}
	flag, result := model.EvaluateProject(project.ID, project.Score)
	if flag {
		ctx.Status(200)
	} else {
		ctx.String(400, result)
	}
}

// QueryProjectsByUser handle
func QueryProjectsByUser(ctx *gin.Context) {
	result, err := model.QueryProjectsByUsername(ctx.Query("username"))
	if err != nil {
		ctx.Status(400)
	} else {
		ctx.JSON(200, result)
	}
}

// QueryProjectByID handle
func QueryProjectByID(ctx *gin.Context) {
	result, err := model.QueryProjectByID(ctx.Query("id"))
	if err != nil {
		ctx.Status(400)
	} else {
		ctx.JSON(200, result)
	}
}
