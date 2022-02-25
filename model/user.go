package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ID       int64
	Email    string `gorm:"type:varchar(50) not null"`
	Name     string `gorm:"type:varchar(20) not null"`
	NickName string `gorm:"type:varchar(20) not null"`
	Password string `gorm:"type:varchar(100) not null"`
	Avatar   string `gorm:"type:varchar(100)"`
}
