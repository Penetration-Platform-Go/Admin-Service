package controller

import (
	"github.com/Penetration-Platform-Go/Admin-Service/model"
	"github.com/gin-gonic/gin"
	"time"
)

// GetInfo handle return user and project
func GetInfo(ctx *gin.Context) {
	var info model.Info
	users, err := model.QueryAllUsers()
	if err != nil {
		ctx.Status(400)
		return
	}
	project, err := model.QueryAllProjects()
	if err != nil {
		ctx.Status(400)
		return
	}
	NotRatedProjectNumber := 0
	for _, each := range project {
		if each.Score == 0 {
			NotRatedProjectNumber++
		}
	}
	temp := []int32{0, 0, 0, 0, 0, 0, 0}
	views, err := model.QueryViews()
	for _, each := range views {
		if each.Date == "all" {
			info.AllViews = each.Number
		} else if each.Date == time.Now().Format("20000101") {
			temp[6] = each.Number
		} else if each.Date == time.Now().AddDate(0, 0, -1).Format("20000101") {
			temp[5] = each.Number
		} else if each.Date == time.Now().AddDate(0, 0, -2).Format("20000101") {
			temp[4] = each.Number
		} else if each.Date == time.Now().AddDate(0, 0, -3).Format("20000101") {
			temp[3] = each.Number
		} else if each.Date == time.Now().AddDate(0, 0, -4).Format("20000101") {
			temp[2] = each.Number
		} else if each.Date == time.Now().AddDate(0, 0, -5).Format("20000101") {
			temp[1] = each.Number
		} else if each.Date == time.Now().AddDate(0, 0, -6).Format("20000101") {
			temp[0] = each.Number
		}
	}
	info.ViewsBeforeWeek = temp
	info.UserNumber = len(users)
	info.AllProjectNumber = len(project)
	info.NotRatedProjectNumber = NotRatedProjectNumber
	ctx.JSON(200, info)
}
