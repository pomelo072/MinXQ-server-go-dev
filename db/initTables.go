package db

import (
	"MinXQ-server-go-dev/models"
	"gorm.io/gorm"
)

func Createtable(db *gorm.DB) {

	db.AutoMigrate(
		&models.User{},
	)
}
