package controllers

import (
	"MinXQ-server-go-dev/config"
	"MinXQ-server-go-dev/database"
	"MinXQ-server-go-dev/models"
	"MinXQ-server-go-dev/utils"
	"fmt"
	"github.com/kataras/iris/v12"
)

// 管理页面进入
func Admin(ctx iris.Context) {
	ck := ctx.GetCookie("usertoken")
	if ck == config.Sysconfig.UserToken {
		ctx.View("admin.html")
	} else {
		ctx.Redirect("/adm/login")
	}
}

// 登录管理页面进入
func AdminLogin(ctx iris.Context) {
	ck := ctx.GetCookie("usertoken")
	if ck == config.Sysconfig.UserToken {
		ctx.Redirect("/adm")
	} else {
		ctx.View("login.html")
	}
}

// 登录验证
func LoginAuth(ctx iris.Context) {
	username := ctx.URLParam("adminName")
	userpsd := ctx.URLParam("adminPsd")
	value := utils.GetSHAEncode(utils.GetSHAEncode(username + userpsd))
	// 测试输出
	fmt.Println(value)
	ctx.SetCookieKV("usertoken", value)
	if value == config.Sysconfig.UserToken {
		ctx.WriteString("200")
	} else {
		ctx.WriteString("400")
	}
}

// 获取统计信息
func Total(ctx iris.Context) {
	// 获取用户总数
	var user models.User
	usertotal := database.Db.Find(&user).RowsAffected
	// 获取点星总数
	var star []int
	database.Db.Table("stars").Pluck("star", &star)
	startotal := 0
	for _, v := range star {
		startotal += v
	}
	totallist := models.Total{User: usertotal, Star: startotal}
	result := utils.GetReturnData(totallist, "SUCCESS")
	ctx.JSON(result)
}
