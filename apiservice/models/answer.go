package models

import "time"

type Answer struct {
	ID            int64
	AnswerDesc    string
	UpVoteBy      string
	DownVoteBy    string
	CountUpVote   int64
	CountDownVote int64
	CreatedBy     string
	CreatedDate   time.Time
}

type AnswerCreateParam struct {
	QuestionID int64  `form:"question_id" json:"question_id" binding:"required"`
	AnswerDesc string `form:"answer_desc" json:"answer_desc" binding:"required"`
	CreatedBy  string `form:"created_by" json:"created_by" binding:"required"`
}

type AnswerUpdateParam struct {
	AnswerDesc string `form:"answer_desc" json:"answer_desc" binding:"required"`
}

type AddUpVoteAnswerParam struct {
	UpVoteBy string `form:"up_vote_by" json:"up_vote_by" binding:"required"`
}

type AddDownVoteAnswerParam struct {
	DownVoteBy string `form:"down_vote_by" json:"down_vote_by" binding:"required"`
}

type QuestionAnswerParam struct {
	QuestionID int64 `form:"question_id" json:"question_id" binding:"required"`
}
