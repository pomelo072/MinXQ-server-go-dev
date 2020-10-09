package controllers

import (
	"MinXQ-server-go-dev/config"
	"MinXQ-server-go-dev/utils"
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
	//fmt.Println(value)
	ctx.SetCookieKV("usertoken", value)
	if value == config.Sysconfig.UserToken {
		ctx.WriteString("200")
	} else {
		ctx.WriteString("400")
	}
}

// 获取统计信息
func Total(ctx iris.Context) {

}
