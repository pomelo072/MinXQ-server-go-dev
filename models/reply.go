package models

// 设计的第一版留言回复表 10.19

// 留言表
type Reply struct {
	MSGID     int    `gorm:"primaryKey"`
	REPLYMSG  string `gorm:"type:varchar(10000)"`
	REPLYNAME string `gorm:"type:varchar(20)"`
	REPLYTIME string `gorm:"type:datetime"`
	REPLYWELL int    `gorm:"type:bigint;default:1"`
}

type Review struct {
	MSGID     int    `gorm:"primaryKey"`
	REPLYMSG  string `gorm:"type:varchar(10000)"`
	REPLYNAME string `gorm:"type:varchar(20)"`
	REPLYTIME int    `gorm:"type:datetime"`
}
