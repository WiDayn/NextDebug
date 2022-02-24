package dto

import "prmlk.com/nextdebug/model"

type ProblemDto struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

//暂时没用上
//func ToProblemDto(problem model.Problem) ProblemDto {
//	return ProblemDto{
//		ID: problem.ID,
//		Name: problem.Name,
//	}
//}

func ToProblemsDto(problem []*model.Problem) []*ProblemDto {
	var info []*ProblemDto
	for i, set := range problem {
		info = append(info, &ProblemDto{0, ""})
		info[i].ID = set.ID
		info[i].Name = set.Name
	}
	return info
}
