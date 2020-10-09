package models

// 星星数据表结构
type Stars struct {
	ID      int    `gorm:"primary_key"`
	ADDRESS string `gorm:"type:varchar(20);index"`
	STAR    int    `gorm:"type:int(10000000)"`
}
