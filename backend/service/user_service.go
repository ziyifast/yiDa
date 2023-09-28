package service

import (
	"github.com/aobco/log"
	"ziyi.com/yiDa/dao"
	"ziyi.com/yiDa/model"
	"ziyi.com/yiDa/pg"
)

type userService struct {
}

var UserService = new(userService)

func (s *userService) AddUser(user *model.User) error {
	defer func() {
		if err := recover(); err != nil {
			log.Errorf("%v", err)
		}
	}()
	//TODO 校验操作
	_, err := dao.UserDao.Insert(pg.Engine, user)
	return err
}
