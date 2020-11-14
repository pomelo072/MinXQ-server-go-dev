package controllers

import (
	"MinXQ-server-go-dev/myhandler"
	"MinXQ-server-go-dev/utils"
	"github.com/kataras/iris/v12"
)

// 点星星
func StarsLight(ctx iris.Context) {
	userid := ctx.URLParam("user_id")
	address := ctx.URLParam("address")
	flag := ctx.URLParam("flag")
	msg := myhandler.Starlight(userid, address, flag)
	result := utils.GetReturnData(msg, "SUCCESS")
	ctx.JSON(result)
}

// 获取地区点星数
func StarsList(ctx iris.Context) {
	list := myhandler.Starlist()
	result := utils.GetReturnData(list, "SUCCESS")
	ctx.JSON(result)
}

// 获取国外数据
func NationList(ctx iris.Context) {
	list := myhandler.Nationlist()
	result := utils.GetReturnData(list, "SUCCESS")
	ctx.JSON(result)
}
