package myhandler

import (
	"MinXQ-server-go-dev/database"
	"MinXQ-server-go-dev/models"
	"time"
)

func Starlight(userid string, address string) string {
	user := new(models.User)
	database.Db.Model(&user).Where("user_id = ?", userid).First(&user)
	if user.LASTSTAR == time.Now().Format("2006-01-02") {
		return "你今天已经点过了噢"
	} else {
		database.Db.Model(&user).Where("user_id = ?", userid).Update("laststar", time.Now().Format("2006-01-02"))
		star := new(models.Stars)
		database.Db.Where("address = ?", address).First(&star)
		star.STAR += 1
		database.Db.Save(&star)
		return "点星成功"
	}
}

func Starlist() interface{} {
	var list []map[string]interface{}
	database.Db.Table("stars").Order("star DESC").Limit(10).Find(&list)
	return list
}
