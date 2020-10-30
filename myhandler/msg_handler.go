package myhandler

import (
	"MinXQ-server-go-dev/database"
	"MinXQ-server-go-dev/models"
	"fmt"
	"gorm.io/gorm"
	"strconv"
	"time"
)

func Addmsg(reply *models.Reply) string {
	msg := reply.REPLYMSG
	msgResult := UseShield(msg)
	if msgResult == "block" {
		return "包含敏感词"
	} else if msgResult == "review" {

	} else if msgResult == "pass" {
		nt := time.Now()
		t := new(models.User)
		database.Db.Table("users").Where("name = ?", reply.REPLYNAME).First(&t)
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

func DelMsg(del *models.Reply) string {
	database.Db.Table("replies").Where("replyname = ? AND replytime = ?", del.REPLYNAME, del.REPLYTIME).Delete(&del)
	return "删除成功"
}

func WellMsg(MsgID string, Username string) string {
	user := new(models.User)
	database.Db.Table("users").Where("name = ?", Username).First(&user)
	lastwell, _ := time.ParseInLocation("2006-01-02 15:04:05", user.LASTWELL, time.Local)
	lastwellid := user.LASTWELLID
	t := time.Now()
	sub := t.Sub(lastwell)
	if sub.Seconds() < 10 {
		return "点太快啦, 10秒后再试噢"
	} else if MsgID == lastwellid {
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

func ListMsg(tp string, pages string, pagesize string) interface{} {
	var list []map[string]interface{}
	if tp == "time" {
		database.Db.Table("replies").Order("replytime DESC").Scopes(Paginate(pages, pagesize)).Find(&list)
		return list
	} else if tp == "well" {
		database.Db.Table("replies").Order("replywell DESC").Scopes(Paginate(pages, pagesize)).Find(&list)
		return list
	} else {
		return "排序类型错误"
	}
}

func Paginate(pages string, pagesizes string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		page, _ := strconv.Atoi(pages)
		if page == 0 {
			page = 1
		}
		pagesize, _ := strconv.Atoi(pagesizes)
		switch {
		case pagesize <= 30:
			pagesize = 30
		case pagesize < 60:
			pagesize = 45
		case pagesize >= 60:
			pagesize = 60
		}

		offset := (page - 1) * pagesize
		return db.Offset(offset).Limit(pagesize)
	}
}
