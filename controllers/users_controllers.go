package controllers

import (
	"MinXQ-server-go-dev/models"
	"MinXQ-server-go-dev/myhandler"
	"MinXQ-server-go-dev/utils"
	"MinXQ-server-go-dev/wxlogin"
	"github.com/kataras/iris/v12"
)

// 登录
func Login(ctx iris.Context) {
	// 获取code
	code := ctx.URLParam("code")
	// 获取openID和session_key
	wxLoginRsp, err := wxlogin.WXLogin(code)
	// 最好别出错, 错了告诉zyx
	if err != nil {
		result := utils.GetReturnData(nil, "FAILED")
		ctx.JSON(result)
	}
	USERID := utils.GetSHAEncode(wxLoginRsp.OpenID)
	// 验证用户是否第一次登录
	user := myhandler.VerifyUserIdExist(USERID)
	if user.USERID == "" {
		// 创建用户记录
		userinfo := myhandler.CreateUser(USERID)
		result := utils.GetReturnData(userinfo, "SUCCESS")
		ctx.JSON(result)
	} else {
		result := utils.GetReturnData(user, "SUCCESS")
		ctx.JSON(result)
	}
}

// 个人信息变更
func PersonalEdit(ctx iris.Context) {
	new_user := new(models.User)
	ctx.ReadJSON(new_user)
	list := myhandler.Personaledit(new_user)
	result := utils.GetReturnData(list, "SUCCESS")
	ctx.JSON(result)
}

// 获取个人信息
func PersonalInfo(ctx iris.Context) {
	userid := ctx.URLParam("user_id")
	if userid == "all" {
		list := myhandler.GetAllpersonalInfo()
		result := utils.GetReturnData(list, "SUCCESS")
		ctx.JSON(result)
	} else {
		list := myhandler.GetpersonalInfo(userid)
		result := utils.GetReturnData(list, "SUCCESS")
		ctx.JSON(result)
	}
}
