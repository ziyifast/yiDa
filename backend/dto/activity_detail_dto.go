package dto

import "ziyi.com/yiDa/common/types"

type ActivityDetailDto struct {
	Id          int             `json:"id"`
	Theme       string          `json:"theme"`
	Owner       int             `json:"owner"`
	Description string          `json:"description"`
	Status      int             `json:"status"`
	UserList    string          `json:"user_list"`
	CreateTime  types.LocalTime `json:"create_time" xorm:"created"`
	UpdateTime  types.LocalTime `json:"update_time" xorm:"updated"`
}
