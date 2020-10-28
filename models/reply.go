package models

// 设计的第一版留言回复表 10.19

// 留言表
type Reply struct {
	MSGID     int    `gorm:"primaryKey"`
	REPLYMSG  string `gorm:"type:varchar(10000)"`
	REPLYNAME string `gorm:"type:varchar(20)"`
	REPLYTIME string `gorm:"type:varchar(15)"`
	REPLYWELL string `gorm:"type:bigint"`
}
