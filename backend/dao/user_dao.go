package dao

import (
	"github.com/aobco/log"
	"xorm.io/xorm"
	"ziyi.com/yiDa/model"
)

type userDao struct {
}

var UserDao = new(userDao)

func (d *userDao) Insert(engine *xorm.Engine, user *model.User) (int64, error) {
	i, err := engine.Insert(user)
	if err != nil {
		log.Errorf("%v", err)
		return -1, err
	}
	return i, nil
}
