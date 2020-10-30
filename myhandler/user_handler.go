package myhandler

import (
	"MinXQ-server-go-dev/database"
	"MinXQ-server-go-dev/models"
)

// 验证用户是否存在
func VerifyUserIdExist(id string) models.User {
	user := models.User{}
	// 根据USERID查询用户
	database.Db.Where("user_id = ?", id).First(&user)
	// 返回User
	return user
}

func CreateUser(id string) models.User {
	user := models.User{USERID: id}
	// 根据USERID创建用户数据记录
	database.Db.Create(&user)
	return user
}

func Personaledit(nuser *models.User) string {
	nName := nuser.NAME
	nameResult := UseShield(nName)
	if nameResult == "pass" {
		err := database.Db.Table("users").Where("user_id = ?", nuser.USERID).Updates(models.User{NAME: nuser.NAME, SCUECID: nuser.SCUECID, COLLEGE: nuser.COLLEGE, ADDRESS: nuser.ADDRESS}).RowsAffected
		if err > 0 {
			return "修改成功"
		} else {
			return "修改失败"
		}
	} else {
		return "包含敏感词"
	}
}

func GetpersonalInfo(userid string) *models.User {
	result := new(models.User)
	err := database.Db.Model(&result).Where("user_id = ?", userid).First(&result).RowsAffected
	if err > 0 {
		return result
	} else {
		return nil
	}
}
