package model

import "gorm.io/gorm"

type OnlineJudge struct {
	gorm.Model
	ID   int64
	Name string `json:"name" gorm:"type:varchar(100) not null"`
	Link string `json:"link" gorm:"type:varchar(200)"`
}

type OnlineJudgeSet struct {
	OnlineJudges *OnlineJudge `json:"online_judges"`
}
