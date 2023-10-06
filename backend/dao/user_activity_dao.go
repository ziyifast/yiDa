package dao

import (
	"xorm.io/xorm"
	"ziyi.com/yiDa/model"
)

type userActivityDao struct{}

var UserActivityDao = new(userActivityDao)

func (d *userActivityDao) GetUserActivityByUidAndAid(engine xorm.Engine, uid, aid int64) (int64, error) {
	i, err := engine.Where("uid = ? and aid = ?", uid, aid).Count()
	if err != nil {
		return 0, err
	}
	return i, nil
}

func (d *userActivityDao) GetUserActivityByUid(engine xorm.Interface, uid int64) (userActivities []*model.UserActivity, err error) {
	err = engine.Where("uid = ?", uid).Find(&userActivities)
	return
}

func (d *userActivityDao) DeleteByUidAndAid(engine xorm.Interface, uid, aid int64) error {
	_, err := engine.Where("uid = ?", uid).Where("aid=?", aid).Delete(&model.UserActivity{})
	if err != nil {
		return err
	}
	return nil
}
