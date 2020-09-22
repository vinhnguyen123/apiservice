package models

import "time"

type Question struct {
	ID            int64
	QuestionDesc  string
	UpVoteBy      string
	DownVoteBy    string
	CountUpVote   int64
	CountDownVote int64
	CreatedBy     string
	CreatedDate   time.Time
}

type QuestionCreateParam struct {
	QuestionDesc string `form:"question_desc" json:"question_desc" binding:"required"`
	CreatedBy    string `form:"created_by" json:"created_by" binding:"required"`
	Tag          string `form:"tag" json:"tag"`
}

type QuestionUpdateParam struct {
	QuestionDesc string `form:"question_desc" json:"question_desc"`
	CreatedBy    string `form:"created_by" json:"created_by"`
}

type AddUpVoteQuestionParam struct {
	UpVoteBy string `form:"up_vote_by" json:"up_vote_by" binding:"required"`
}

type AddDownVoteQuestionParam struct {
	DownVoteBy string `form:"down_vote_by" json:"down_vote_by" binding:"required"`
}

type QuestionListParam struct {
	Limit  int `form:"limit" json:"limit" binding:"required"`
	Offset int `form:"offset" json:"offset"`
}

type GetQuestionTag struct {
	ID            int64
	QuestionDesc  string
	UpVoteBy      string
	DownVoteBy    string
	CountUpVote   int64
	CountDownVote int64
	CreatedBy     string
	CreatedDate   time.Time
	TagDesc       string `gorm:"column:tag_desc"`
}

func (gu *GetQuestionTag) TableName() string {
	return "questions"
}
