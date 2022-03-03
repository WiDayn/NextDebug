package model

import "gorm.io/gorm"

type TestSet struct {
	gorm.Model
	ID        int64
	ProblemID int    `json:"problem_id" gorm:"type:int(11) not null"`
	Uploader  int    `json:"uploader" gorm:"type:int(11) not null"`
	Votes     int    `json:"votes" gorm:"type:int(11) not null"`
	Input     string `json:"input" gorm:"type:text"`
	Output    string `json:"output" gorm:"type:text"`
}
