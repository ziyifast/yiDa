package model

import (
	"ziyi.com/yiDa/common/types"
)

type ActivityDetail struct {
	Id          int             `json:"id"`
	Theme       string          `json:"theme"`
	Owner       int             `json:"owner"`
	Description string          `json:"description"`
	Status      int             `json:"status"`
	CreateTime  types.LocalTime `json:"create_time" xorm:"created"`
	UpdateTime  types.LocalTime `json:"update_time" xorm:"updated"`
}

func (a *ActivityDetail) TableName() string {
	return "activity_detail"
}
