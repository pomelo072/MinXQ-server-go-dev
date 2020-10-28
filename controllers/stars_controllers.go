package controllers

import (
	"MinXQ-server-go-dev/myhandler"
	"MinXQ-server-go-dev/utils"
	"github.com/kataras/iris/v12"
)

// 点星星
func StarsLight(ctx iris.Context) {
	userid := ctx.URLParam("user_id")
	addess := ctx.URLParam("address")
	msg := myhandler.Starlight(userid, addess)
	result := utils.GetReturnData(msg, "SUCCESS")
	ctx.JSON(result)
}

// 获取排行榜
func StarsList(ctx iris.Context) {
	list := myhandler.Starlist()
	result := utils.GetReturnData(list, "SUCCESS")
	ctx.JSON(result)
}
