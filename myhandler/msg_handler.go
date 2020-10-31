package myhandler

import (
	"MinXQ-server-go-dev/database"
	"MinXQ-server-go-dev/models"
	"MinXQ-server-go-dev/utils"
	"fmt"
	"strconv"
	"time"
)

// 留言提交
func Addmsg(reply *models.Reply) string {
	msg := reply.REPLYMSG
	// 屏蔽词检测
	msgResult := UseShield(msg)
	// 检测到敏感词
	if msgResult == "block" {
		return "包含敏感词"
	} else if msgResult == "review" {
		// 可能存在敏感词, 进入审核表
		review := models.Review{REPLYMSG: reply.REPLYMSG, REPLYNAME: reply.REPLYNAME, REPLYTIME: time.Now().Format("2006-01-02 15:04:05")}
		database.Db.Table("reviews").Create(&review)
		return "留言中可能存在敏感词, 待人工审核通过过后就会发布"
	} else if msgResult == "pass" {
		// 不存在敏感词
		nt := time.Now()
		t := new(models.User)
		database.Db.Table("users").Where("name = ?", reply.REPLYNAME).First(&t)
		// 留言间隔时间判定 > 10分钟
		ft, _ := time.ParseInLocation("2006-01-02 15:04:05", t.LASTREPLY, time.Local)
		sub := nt.Sub(ft)
		if sub.Minutes() >= 10 {
			reply.REPLYTIME = nt.Format("2006-01-02 15:04:05")
			database.Db.Select("replymsg", "replyname", "replytime").Create(&reply)
			database.Db.Model(&models.User{}).Where("name = ?", reply.REPLYNAME).Update("lastreply", reply.REPLYTIME)
			return "留言成功"
		} else {
			return fmt.Sprintf("留言时间要大于10分钟哦, 你的上次留言时间是%s", t.LASTREPLY)
		}
	}
	return msgResult
}

// 删除留言, 回复人, 回复时间
func DelMsg(del *models.Reply) string {
	database.Db.Table("replies").Where("replyname = ? AND replytime = ?", del.REPLYNAME, del.REPLYTIME).Delete(&del)
	return "删除成功"
}

// 点赞留言
func WellMsg(MsgID string, Username string) string {
	user := new(models.User)
	database.Db.Table("users").Where("name = ?", Username).First(&user)
	// 查找用户的点赞时间和上条点赞
	lastwell, _ := time.ParseInLocation("2006-01-02 15:04:05", user.LASTWELL, time.Local)
	lastwellid := user.LASTWELLID
	t := time.Now()
	sub := t.Sub(lastwell)
	if sub.Seconds() < 10 {
		// 判断点赞间隔
		return "点太快啦, 10秒后再试噢"
	} else if MsgID == lastwellid {
		// 判断是否点赞
		return "你已经给它点过啦, 去看看别的吧"
	} else {
		msg := new(models.Reply)
		database.Db.Table("replies").Where("msg_id = ?", MsgID).First(&msg)
		msg.REPLYWELL += 1
		database.Db.Table("replies").Where("msg_id = ?", MsgID).Update("replywell", msg.REPLYWELL)
		database.Db.Table("users").Where("user_id = ?", user.USERID).Updates(&models.User{LASTWELL: t.Format("2006-01-02 15:04:05"), LASTWELLID: MsgID})
		return "点赞成功"
	}
}

// 获取留言
func ListMsg(tp string, pages string, pagesize string) interface{} {
	var list []map[string]interface{}
	if tp == "time" {
		database.Db.Table("replies").Order("replytime DESC").Scopes(utils.Paginate(pages, pagesize)).Find(&list)
		return list
	} else if tp == "well" {
		database.Db.Table("replies").Order("replywell DESC").Scopes(utils.Paginate(pages, pagesize)).Find(&list)
		return list
	} else {
		return "排序类型错误"
	}
}

// 获取待审核留言 (后台)
func ListReview() interface{} {
	var list []map[string]interface{}
	database.Db.Table("reviews").Find(&list)
	return list
}

// 通过审核留言
func PassReview(msgid string) string {
	id, _ := strconv.Atoi(msgid)
	reply := new(models.Reply)
	database.Db.Table("reviews").Where("msg_id = ?", id).First(&reply)
	database.Db.Table("replies").Select("replymsg", "replyname", "replytime").Create(&reply)
	database.Db.Table("reviews").Where("msg_id = ?", id).Delete(&reply)
	return "审核通过, 操作成功"
}

// 不通过审核留言
func DelReview(msgid string) string {
	id, _ := strconv.Atoi(msgid)
	review := new(models.Review)
	database.Db.Table("reviews").Where("msg_id = ?", id).Delete(&review)
	return "审核不通过, 操作成功"
}
