package database

import (
	"MinXQ-server-go-dev/models"
)

func Createtable() {
	// 数据库自动迁移
	Db.AutoMigrate(
		&models.User{},
		&models.Stars{},
		&models.Reply{},
		&models.Review{},
	)
	//Db.Model(&models.Stars{}).Create([]map[string]interface{}{
	//	{"address": "HN1"}, // 湖南
	//	{"address": "GX"},  // 广西
	//	{"address": "HN2"}, // 河南
	//	{"address": "HN3"}, // 海南
	//	{"address": "BJ"},  // 北京
	//	{"address": "TJ"},  // 天津
	//	{"address": "JL"},  // 吉林
	//	{"address": "SH"},  // 上海
	//	{"address": "GD"},  // 广东
	//	{"address": "SC"},  // 四川
	//	{"address": "HB1"}, // 湖北
	//	{"address": "HB2"}, // 河北
	//	{"address": "HK"},  // 香港
	//	{"address": "FJ"},  // 福建
	//})
}
