package model

import "time"

type ActivityDetail struct {
	Id          int
	Theme       string
	Owner       string
	Description string
	UserList    string
	Status      int
	CreateTime  time.Time
	UpdateTime  time.Time
}

func (a *ActivityDetail) TableName() string {
	return "activity_detail"
}
