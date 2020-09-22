package models

type Tag struct {
	ID      int64
	TagDesc string
}

type TagParam struct {
	TagDesc string `form:"tag" json:"tag" binding:"required"`
}
