package dto

import (
	"prmlk.com/nextdebug/common"
	"prmlk.com/nextdebug/model"
	"time"
)

type TestSetDto struct {
	ID         int64
	CreatedAt  string `json:"created_at"`
	Uploader   string `json:"uploader"`
	UploaderID int    `json:"uploader_id"`
	Votes      int    `json:"votes"`
	Input      string `json:"input"`
	Output     string `json:"output"`
}

func ToTestSetsDto(TestSet []*model.TestSet) []*TestSetDto {
	db := common.GetDB()
	var info []*TestSetDto
	for i, set := range TestSet {
		var uploader model.User
		db.Model(model.User{}).Where("ID = ?", set.Uploader).First(&uploader)
		info = append(info, &TestSetDto{0, time.Now().Format("2006-01-02"), "", 0, 0, "", ""})
		info[i].ID = set.ID
		info[i].CreatedAt = set.CreatedAt.Format("2006-01-02")
		info[i].Uploader = uploader.NickName
		info[i].UploaderID = set.Uploader
		info[i].Votes = set.Votes
		info[i].Input = set.Input
		info[i].Output = set.Output
	}
	return info
}
