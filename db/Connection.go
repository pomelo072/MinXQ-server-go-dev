package db

import (
	"MinXQ-server-go-dev/config"
	"gorm.io/driver/mysql"
	_ "gorm.io/driver/mysql"
	"gorm.io/gorm"
	"strings"
)

func init() {
	dsn := strings.Join([]string{config.Sysconfig.DBUserName, ":", config.Sysconfig.DBPassword, "@(", config.Sysconfig.DBIp, ":", config.Sysconfig.DBPort, ")/", config.Sysconfig.DBName, "?charset=utf8mb4&parseTime=true&loc=Local"}, "")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
}
