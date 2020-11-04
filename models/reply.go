package models

// 留言表
type Reply struct {
	MSGID     int    `gorm:"primaryKey"`
	REPLYMSG  string `gorm:"type:varchar(10000)"`
	REPLYNAME string `gorm:"type:varchar(20)"`
	USERID    string `gorm:"type:varchar(70);index"`
	REPLYTIME string `gorm:"type:varchar(30);default:"2006-01-02 15:04:05""`
	REPLYWELL int    `gorm:"type:bigint;default:1"`
}

type Review struct {
	MSGID     int    `gorm:"primaryKey"`
	REPLYMSG  string `gorm:"type:varchar(10000)"`
	REPLYNAME string `gorm:"type:varchar(20)"`
	USERID    string `gorm:"type:varchar(70);index"`
	REPLYTIME string `gorm:"type:varchar(30);default:"2006-01-02 15:04:05""`
}
