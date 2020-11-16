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
	nt := time.Now()
	t := new(models.User)
	database.Db.Table("users").Where("user_id = ?", reply.USERID).First(&t)
	// 留言间隔时间判定 > 10分钟
	ft, _ := time.ParseInLocation("2006-01-02 15:04:05", t.LASTREPLY, time.Local)
	sub := nt.Sub(ft)
	if sub.Minutes() >= 10 {
		// 检测到敏感词
		if msgResult == "block" {
			return "包含敏感词"
		} else {
			database.Db.Model(&models.User{}).Where("user_id = ?", reply.USERID).Update("lastreply", nt.Format("2006-01-02 15:04:05"))
			u := new(models.User)
			database.Db.Table("users").Where("user_id = ?", reply.USERID).First(&u)
			review := models.Review{USERID: reply.USERID, REPLYMSG: reply.REPLYMSG, REPLYNAME: reply.REPLYNAME, COLLEGE: u.COLLEGE, REPLYTIME: nt.Format("2006-01-02 15:04:05")}
			database.Db.Table("reviews").Create(&review)
			return "留言成功, 待人工审核通过过后就会发布"
		}
	} else {
		return fmt.Sprintf("留言时间要大于10分钟哦, 你的上次留言时间是%s", t.LASTREPLY)
	}
	return msgResult
}

// 删除留言

//func DelMsg(id, time string) string {
//	del := new(models.Reply)
//	database.Db.Table("replies").Where("user_id = ? AND replytime = ?", id, time).Delete(&del)
//	return "删除成功"
//}

func DelMsg(id string) string {
	del := new(models.Reply)
	database.Db.Table("replies").Where("msg_id = ?", id).Delete(&del)
	return "删除成功"
}

// 点赞留言
func WellMsg(MsgID string, Userid string) string {
	msg := new(models.Reply)
	database.Db.Table("replies").Where("msg_id = ?", MsgID).First(&msg)
	msg.REPLYWELL += 1
	database.Db.Table("replies").Where("msg_id = ?", MsgID).Update("replywell", msg.REPLYWELL)
	msgid, _ := strconv.Atoi(MsgID)
	well := models.Well{USERID: Userid, MSGID: msgid}
	database.Db.Table("wells").Create(&well)
	return "点赞成功"
}

// 用户点赞列表
func Userwell(Userid string) []int64 {
	var well []int64
	welllist := database.Db.Table("wells").Where("user_id = ?", Userid)
	welllist.Pluck("msg_id", &well)
	return well
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
	database.Db.Table("replies").Select("replymsg", "replyname", "replytime", "user_id", "college").Create(&reply)
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
