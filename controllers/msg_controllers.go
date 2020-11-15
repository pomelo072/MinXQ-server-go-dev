package controllers

import (
	"MinXQ-server-go-dev/models"
	"MinXQ-server-go-dev/myhandler"
	"MinXQ-server-go-dev/utils"
	"github.com/kataras/iris/v12"
)

// 增加留言
func MsgAdd(ctx iris.Context) {
	reply := new(models.Reply)
	ctx.ReadJSON(&reply)
	list := myhandler.Addmsg(reply)
	result := utils.GetReturnData(list, "SUCCESS")
	ctx.JSON(result)
}

// 点赞留言
func MsgWell(ctx iris.Context) {
	MsgID := ctx.URLParam("msgid")
	Userid := ctx.URLParam("userid")
	list := myhandler.WellMsg(MsgID, Userid)
	result := utils.GetReturnData(list, "SUCCESS")
	ctx.JSON(result)
}

// 删除留言
func MsgDelete(ctx iris.Context) {
	//userid := ctx.URLParam("user_id")
	//replytime := ctx.URLParam("replytime")
	//delMsg := myhandler.DelMsg(userid, replytime)
	msgid := ctx.URLParam("msgid")
	delMsg := myhandler.DelMsg(msgid)
	result := utils.GetReturnData(delMsg, "SUCCESS")
	ctx.JSON(result)
}

// 获取留言
func MsgList(ctx iris.Context) {
	tp := ctx.URLParam("type")
	pages := ctx.URLParam("pages")
	pagesize := ctx.URLParam("pagesize")
	list := myhandler.ListMsg(tp, pages, pagesize)
	result := utils.GetReturnData(list, "SUCCESS")
	ctx.JSON(result)
}

// 通过审核
func ReviewPass(ctx iris.Context) {
	msgid := ctx.URLParam("msgid")
	list := myhandler.PassReview(msgid)
	result := utils.GetReturnData(list, "SUCCESS")
	ctx.JSON(result)
}

// 不通过审核
func ReviewDel(ctx iris.Context) {
	msgid := ctx.URLParam("msgid")
	list := myhandler.DelReview(msgid)
	result := utils.GetReturnData(list, "SUCCESS")
	ctx.JSON(result)
}

// 获取待审核留言
func ReviewList(ctx iris.Context) {
	list := myhandler.ListReview()
	result := utils.GetReturnData(list, "SUCCESS")
	ctx.JSON(result)
}
