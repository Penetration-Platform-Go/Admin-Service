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

// DeleteProjectByID handle
func DeleteProjectByID(ctx *gin.Context) {
	flag, result := model.DeleteProjectByID(ctx.Query("id"))
	if flag {
		ctx.Status(200)
	} else {
		ctx.String(400, result)
	}
}

// DeleteProjectByUsername handle
func DeleteProjectByUsername(ctx *gin.Context) {
	flag, result := model.DeleteProjectByUsername(ctx.Query("username"))
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

// QueryProjectByTitle handle
func QueryProjectByTitle(ctx *gin.Context) {
	result, err := model.QueryProjectByTitle(ctx.Query("title"))
	if err != nil {
		ctx.Status(400)
	} else {
		ctx.JSON(200, result)
	}
}
