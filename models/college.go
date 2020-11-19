package models

type CollegeList struct {
	ID          int    `gorm:"primaryKey"`
	COLLEGEID   string `gorm:"type:varchar(3)"`
	COLLEGENAME string `gorm:"type:varchar(15)"`
	STAR        int64  `gorm:"default:1"`
}
