package models

// 用户数据表结构
type User struct {
	ID      int    `gorm:"primary_key"`
	USERID  string `gorm:"type:varchar(70);index"`
	NAME    string `gorm:"type:varchar(20)"`
	SCUECID string `gorm:"type:varchar(15)"`
	COLLEGE string `gorm:"type:varchar(20)"`
	ADDRESS string `gorm:"type:varchar(20)"`
}
