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
		&models.Comment{},
	)
}
