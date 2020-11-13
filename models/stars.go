package models

// 星星数据表结构
type Stars struct {
	ID      int    `gorm:"primaryKey"`
	NATION  string `gorm:"type:varchar(20)"`
	ADDRESS string `gorm:"type:varchar(20);index"`
	STAR    int64  `gorm:"default:1"`
}
