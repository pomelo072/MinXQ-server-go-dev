package db

import (
	"MinXQ-server-go-dev/config"
	"gorm.io/driver/mysql"
	_ "gorm.io/driver/mysql"
	"gorm.io/gorm"
	"strings"
	"time"
)

var db *gorm.DB

func init() {
	dsn := strings.Join([]string{config.Sysconfig.DBUserName, ":", config.Sysconfig.DBPassword, "@(", config.Sysconfig.DBIp, ":", config.Sysconfig.DBPort, ")/", config.Sysconfig.DBName, "?charset=utf8mb4&parseTime=true&loc=Local"}, "")
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       dsn,
		SkipInitializeWithVersion: false,
	}), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	sqlDB, err := db.DB()
	if err != nil {
		panic(err)
	}
	sqlDB.SetMaxIdleConns(20)
	sqlDB.SetMaxOpenConns(2000)
	sqlDB.SetConnMaxLifetime(1 * time.Second)
	Createtable(db)
}
