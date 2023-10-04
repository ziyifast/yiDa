package model

import (
	"time"

	"ziyi.com/yiDa/common/types"
)

type User struct {
	Id          int             `json:"id"`
	Username    string          `json:"username"`
	Password    string          `json:"password"`
	Wechat      string          `json:"wechat"`
	Phone       string          `json:"phone"`
	City        string          `json:"city"`
	BirthDay    time.Time       `json:"birthday" xorm:"birthday"`
	Hobbies     string          `json:"hobbies"`
	Description string          `json:"description"`
	CreateTime  types.LocalTime `json:"create_time" xorm:"create_time timestamp created"`
	UpdateTime  types.LocalTime `json:"update_time" xorm:"update_time timestamp updated"`
}

func (u *User) TableName() string {
	return "user"
}
