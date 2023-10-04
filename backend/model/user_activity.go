package model

type UserActivity struct {
	Uid int64 `json:"uid"`
	Aid int64 `json:"aid"`
}

func (*UserActivity) TableName() string {
	return "user_activity"
}
