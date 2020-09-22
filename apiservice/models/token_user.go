package models

import "time"

type TokenUser struct {
	ID        int64
	UserID    int64
	AccessUid string
	AtExpires time.Time
}
