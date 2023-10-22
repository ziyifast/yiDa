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

func (d *userActivityDao) GetUserActivityDetailByUidAndAid(engine xorm.Engine, uid, aid int64) (bool, error) {
	userActivity := new(model.UserActivity)
	b, err := engine.Where("uid = ? and aid = ?", uid, aid).Get(userActivity)
	if err != nil {
		return false, err
	}
	return b, nil
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

func (d *userActivityDao) InsertUserActivity(engine xorm.Interface, userActivity *model.UserActivity) error {
	_, err := engine.Insert(userActivity)
	if err != nil {
		return err
	}
	return nil
}
