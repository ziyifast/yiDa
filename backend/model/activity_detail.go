package model

import (
	"ziyi.com/yiDa/common/types"
)

type ActivityDetail struct {
	Id          int             `json:"id"`
	Theme       string          `json:"theme"`
	Owner       int             `json:"owner"`
	Description string          `json:"description"`
	UserList    string          `json:"user_list"`
	Status      int             `json:"stats"`
	CreateTime  types.LocalTime `json:"create_time"`
	UpdateTime  types.LocalTime `json:"update_time"`
}

func (a *ActivityDetail) TableName() string {
	return "activity_detail"
}
