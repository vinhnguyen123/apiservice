package models

import "time"

type User struct {
	ID                int64
	UserName          string
	Email             string
	EncryptedPassword string
	Salt              string
	CreatedDate       time.Time
}

type UserParam struct {
	UserName string `form:"user_name" json:"user_name" binding:"required"`
	Email    string `form:"email" json:"email" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

type UpdateUserParam struct {
	UserName string `form:"user_name" json:"user_name" binding:"required"`
	Email    string `form:"email" json:"email" binding:"required"`
}

type UserLogin struct {
	Email    string `form:"email" json:"email" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

type GetUser struct {
	ID          int64
	UserName    string
	Email       string
	CreatedDate time.Time
}

func (gu *GetUser) TableName() string {
	return "users"
}
