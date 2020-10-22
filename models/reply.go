package models

// 设计的第一版留言回复表 10.19

// 留言表
type Reply struct {
	MSGID     int    `gorm:"primaryKey"`
	REPLYMSG  string `gorm:"type:varchar(10000)"`
	REPLYNAME string `gorm:"type:varchar(20)"`
	REPLYTIME string `gorm:"type:datetime"`
	REPLYWELL string `gorm:"type:int"`
}

// 回复表
type Comment struct {
	COMID    int    `gorm:"primaryKey"`
	COMMSG   string `gorm:"type:varchar(10000)"`
	COMNAME  string `gorm:"type:varchar(20)"`
	COMTIME  string `gorm:"type:datetime"`
	MSGID    int    `gorm:"type:int"`
	LEVEL    int    `gorm:"type:tinyint"`
	FATHERID int    `gorm:"type:int"`
}
