package myhandler

import (
	"MinXQ-server-go-dev/database"
	"MinXQ-server-go-dev/models"
	"time"
)

type Nation struct {
	nation string
	star   int
	data   interface{}
}

type Nationls struct {
	china Nation
	other interface{}
}

// 点星触发, 只能点一次
func Starlight(userid string, address string) string {
	user := new(models.User)
	database.Db.Model(&user).Where("user_id = ?", userid).First(&user)
	if user.LASTSTAR == time.Now().Format("2006-01-02") {
		return "你今天已经点过了噢"
	} else {
		database.Db.Model(&user).Where("user_id = ?", userid).Update("laststar", time.Now().Format("2006-01-02"))
		star := new(models.Stars)
		database.Db.Where("address = ?", address).FirstOrCreate(&star, models.Stars{ADDRESS: address})
		star.STAR += 1
		database.Db.Save(&star)
		return "点星成功"
	}
}

// 获取排行榜
func Starlist() interface{} {
	var list []map[string]interface{}
	database.Db.Table("stars").Order("star DESC").Find(&list)
	return list
}

func Nationlist() interface{} {
	var Chinalist []map[string]interface{}
	var Otherlist []map[string]interface{}
	database.Db.Table("stars").Where("nation = ?", "中国").Order("star DESC").Limit(10).Find(&Chinalist)
	database.Db.Table("stars").Not("nation = ?", "中国").Order("star DESC").Find(&Otherlist)
	var chinastar []int
	database.Db.Table("stars").Where("nation = ?", "中国").Pluck("star", &chinastar)
	chinatotal := 0
	for _, v := range chinastar {
		chinatotal += v
	}
	china := Nation{"中国", chinatotal, Chinalist}
	result := Nationls{china: china, other: Otherlist}
	return result
}
