package model

import "gorm.io/gorm"

type Problem struct {
	gorm.Model
	ID             int64
	Name           string        `json:"name" gorm:"type:varchar(100) not null"`
	OriginalID     string        `json:"original_id" gorm:"type:varchar(100)"`
	Description    string        `json:"description" gorm:"type:text"`
	From           int           `json:"from" gorm:"type:int(11)"`
	ProblemLink    string        `json:"problem_link" gorm:"type:varchar(200)"`
	Uploader       int           `json:"uploader" gorm:"type:int(11)"`
	ProblemList    []ProblemList `json:"problem_list" gorm:"many2many:problems_problemLists"`
	ProblemTag     []ProblemTag  `json:"problem_tag" gorm:"many2many:problems_problemTags"`
	RelatedProblem []Problem     `json:"related_problems" gorm:"many2many:related_problems"`
}

type ProblemSet struct {
	Problems *Problem `json:"problems"`
}

type ProblemTag struct {
	gorm.Model
	ID          int64
	Name        string    `json:"name" gorm:"type:varchar(100) not null unique"`
	Description string    `json:"description" gorm:"type:text"`
	Problems    []Problem `json:"problem" gorm:"many2many:problems_problemTags"`
}

type ProblemList struct {
	gorm.Model
	ID       int64
	Name     string    `json:"name" gorm:"type:varchar(100) not null"`
	Creator  int       `json:"Creator" gorm:"type:int(11)"`
	Problems []Problem `json:"problem_list" gorm:"many2many:problems_problemLists"`
	Vote     int64     `json:"vote" gorm:"type:int(11)"`
}
