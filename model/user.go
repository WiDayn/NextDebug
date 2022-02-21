package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ID       int64
	Name     string `gorm:"type:varchar(20); not null"`
	Password string `gorm:"type:varchar(100); not null"`
}
