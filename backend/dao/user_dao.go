package dao

import (
	"errors"

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

func (d *userDao) UpdateUserById(engine *xorm.Engine, user *model.User) (int64, error) {
	i, err := engine.Where("id = ?", user.Id).Update(user)
	if i < 1 {
		return -1, errors.New("the user is not found")
	}
	return i, err
}

func (d *userDao) GetUserById(engine *xorm.Engine, id int64) (*model.User, error) {
	user := new(model.User)
	_, err := engine.Where("id = ?", id).Get(user)
	if err != nil {
		return nil, errors.New("the user is not found")
	}
	return user, nil
}

func (d *userDao) GetUsersByName(engine *xorm.Engine, name string) ([]*model.User, error) {
	users := make([]*model.User, 0)
	err := engine.Where("username like %?", name).Find(&users)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (d *userDao) GetUserByName(engine *xorm.Engine, name string) (*model.User, error) {
	user := new(model.User)
	_, err := engine.Where("username = ?", name).Get(user)
	if err != nil {
		return nil, errors.New("the user is not found")
	}
	return user, nil
}

func (d *userDao) DeleteUserById(engine *xorm.Engine, id int64) (int64, error) {
	i, err := engine.Where("id = ?", id).Delete(&model.User{})
	if err != nil {
		return -1, err
	}
	return i, nil
}
