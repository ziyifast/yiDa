package model

import (
	"ziyi.com/yiDa/common/types"
)

type User struct {
	Id          int             `json:"id"`
	Username    string          `json:"username"`
	Password    string          `json:"password"`
	Wechat      string          `json:"wechat"`
	Phone       string          `json:"phone"`
	City        string          `json:"city"`
	Description string          `json:"description"`
	CreateTime  types.LocalTime `json:"create_time"`
	UpdateTime  types.LocalTime `json:"update_time"`
}

func (u *User) TableName() string {
	return "user"
}
