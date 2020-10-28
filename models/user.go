package models

// 用户数据表结构
type User struct {
	ID       int    `gorm:"primaryKey"`
	USERID   string `gorm:"type:varchar(70);index"`
	NAME     string `gorm:"type:varchar(20)"`
	SCUECID  string `gorm:"type:varchar(12)"`
	COLLEGE  string `gorm:"type:varchar(20)"`
	ADDRESS  string `gorm:"type:varchar(20)"`
	LASTSTAR string `gorm:"type:varchar(15)"`
}

type Users struct {
	User User
}
