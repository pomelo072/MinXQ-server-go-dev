package main

import (
	"MinXQ-server-go-dev/config"
	"MinXQ-server-go-dev/controllers"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/core/router"
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
	app.PartyFunc("/stars", func(stars router.Party) {
		stars.Get("/light", controllers.StarsLight) // 点亮
		stars.Get("/list", controllers.StarsList)   // 排行榜
	})

	// 消息管理
	app.PartyFunc("/message", func(message router.Party) {
		// 屏蔽词管理
		message.PartyFunc("/shield", func(shield router.Party) {
			shield.Post("/add", controllers.ShieldAdd)       // 增加屏蔽词
			shield.Post("/delete", controllers.ShieldDelete) // 删除屏蔽词
			shield.Get("/list", controllers.ShieldList)      // 查看屏蔽词
		})
		// 留言管理
		message.PartyFunc("/msg", func(msg router.Party) {
			msg.Post("/add", controllers.MsgAdd)       // 提交留言or回复
			msg.Post("/well", controllers.MsgWell)     // 点赞
			msg.Post("/delete", controllers.MsgDelete) // 删除留言or回复
			msg.Get("/list", controllers.MsgList)      // 查看留言
			// 表情包部分
			msg.Post("/imgAdd", controllers.ImgAdd)  // 增加表情包
			msg.Post("/imgDel", controllers.ImgDel)  // 删除表情包
			msg.Get("/imgList", controllers.ImgList) // 查看表情包
		})
	})
	// 个人界面
	app.PartyFunc("/personal", func(personal router.Party) {
		personal.Get("/login", controllers.Login)
		personal.Post("/edit", controllers.PersonalEdit) // 修改个人信息
		personal.Get("/info", controllers.PersonalInfo)  // 查看个人信息
	})

	// 统计模块
	app.PartyFunc("/count", func(count router.Party) {
		count.Get("/total", controllers.Total) // 总用户量和总点亮数
	})
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
