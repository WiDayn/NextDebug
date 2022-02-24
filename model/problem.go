package model

import "gorm.io/gorm"

type Problem struct {
	gorm.Model
	ID          int64
	Name        string `json:"name" gorm:"type:varchar(100) not null"`
	Description string `json:"description" gorm:"type:text"`
	From        int    `json:"from" gorm:"type:int(11)"`
	ProblemLink string `json:"problem_link" gorm:"type:varchar(200)"`
	Uploader    int    `json:"uploader" gorm:"type:int(11)"`
}

type ProblemSet struct {
	Problems *Problem `json:"problems"`
}
