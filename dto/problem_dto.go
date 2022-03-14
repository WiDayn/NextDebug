package dto

import (
	"prmlk.com/nextdebug/common"
	"prmlk.com/nextdebug/model"
	"time"
)

type ProblemDto struct {
	ID         int64              `json:"id"`
	Name       string             `json:"name"`
	OriginalID string             `json:"original_id"`
	From       string             `json:"from"`
	ProblemTag []model.ProblemTag `json:"problem_tag"`
}

type ProblemDetailDto struct {
	ID             int64
	CreatedAt      time.Time           `json:"created_at"`
	Name           string              `json:"name"`
	OriginalID     string              `json:"original_id"`
	Description    string              `json:"description"`
	From           string              `json:"from"`
	ProblemLink    string              `json:"problem_link"`
	Uploader       string              `json:"uploader"`
	ProblemList    []model.ProblemList `json:"problem_list"`
	ProblemTag     []model.ProblemTag  `json:"problem_tag"`
	RelatedProblem []model.Problem     `json:"related_problems"`
}

func ToProblemDetailDto(problem model.Problem) ProblemDetailDto {
	db := common.GetDB()
	var fromOJ model.OnlineJudge
	db.Model(model.OnlineJudge{}).Where("ID = ?", problem.From).First(&fromOJ)
	var uploader model.User
	db.Model(model.User{}).Where("ID = ?", problem.Uploader).First(&uploader)
	return ProblemDetailDto{
		ID:             problem.ID,
		CreatedAt:      problem.CreatedAt,
		Name:           problem.Name,
		OriginalID:     problem.OriginalID,
		Description:    problem.Description,
		From:           fromOJ.Name,
		ProblemLink:    problem.ProblemLink,
		Uploader:       uploader.NickName,
		ProblemList:    problem.ProblemList,
		ProblemTag:     problem.ProblemTag,
		RelatedProblem: problem.RelatedProblem,
	}
}

func ToProblemsDto(problem []*model.Problem) []*ProblemDto {
	db := common.GetDB()
	var info []*ProblemDto
	for i, set := range problem {
		var fromOJ model.OnlineJudge
		db.Model(model.OnlineJudge{}).Where("ID = ?", set.From).First(&fromOJ)
		info = append(info, &ProblemDto{0, "", "", "", nil})
		info[i].ID = set.ID
		info[i].Name = set.Name
		info[i].OriginalID = set.OriginalID
		info[i].From = fromOJ.Name
		info[i].ProblemTag = set.ProblemTag
	}
	return info
}
