package models

type Well struct {
	ID     int    `gorm:"primaryKey"`
	USERID string `gorm:"type:varchar(70);index"`
	MSGID  int    `gorm:"type:int"`
}
