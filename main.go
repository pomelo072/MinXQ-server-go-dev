package main

import (
	"MinXQ-server-go-dev/config"
	"MinXQ-server-go-dev/controllers"
	"github.com/kataras/iris/v12"
)

func main() {
	// 创建App
	app := iris.New()
	app.Use(Cors)
	// 注册视图模板
	app.RegisterView(iris.HTML("./views", ".html"))
	// 管理界面登录
	app.Get("/adm", controllers.Admin)
	app.Get("/adm/login", controllers.AdminLogin)
	app.Post("/adm/LoginAuth", controllers.LoginAuth)

	// 路由分组

	// 点亮星星
	//app.Party("/stars", controllers.Stars)

	// 留言管理

	app.Run(iris.Addr(config.Sysconfig.Port))
}

func Cors(ctx iris.Context) {
	ctx.Header("Access-Control-Allow-Origin", "*")
	if ctx.Request().Method == "OPTIONS" {
		ctx.Header("Access-Control-Allow-Methods", "GET,POST,PUT,DELETE,PATCH,OPTIONS")
		ctx.Header("Access-Control-Allow-Headers", "Content-Type, Accept, Authorization")
		ctx.StatusCode(204)
		return
	}
	ctx.Next()
}
