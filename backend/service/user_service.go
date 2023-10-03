package service

import (
	"fmt"

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

func (s *userService) UpdateUser(user *model.User) error {
	defer func() {
		if err := recover(); err != nil {
			log.Errorf("%v", err)
		}
	}()
	u, err := dao.UserDao.GetUserById(pg.Engine, int64(user.Id))
	if err != nil {
		log.Errorf("%v", err)
		return err
	}
	if u == nil {
		return fmt.Errorf("the user is not exist")
	}
	//TODO 校验操作
	if _, err := dao.UserDao.UpdateUserById(pg.Engine, user); err != nil {
		log.Errorf("%v", err)
		return err
	}
	return nil
}

func (s *userService) GetUserById(id int64) (*model.User, error) {
	if id < 0 {
		return nil, fmt.Errorf("the id is invalid")
	}
	user, err := dao.UserDao.GetUserById(pg.Engine, id)
	if err != nil {
		log.Errorf("%v", err)
		return nil, err
	}
	return user, nil
}

func (s *userService) GetUserByName(name string) (*model.User, error) {
	if name == "" {
		return nil, fmt.Errorf("the name can not by empty")
	}
	user, err := dao.UserDao.GetUserByName(pg.Engine, name)
	if err != nil {
		log.Errorf("%v", err)
		return nil, err
	}
	return user, nil
}

func (s *userService) DeleteUserById(id int64) error {
	_, err := dao.UserDao.DeleteUserById(pg.Engine, id)
	if err != nil {
		log.Errorf("%v", err)
		return err
	}
	return nil
}
