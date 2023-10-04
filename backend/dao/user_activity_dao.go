package dao

import (
	"xorm.io/xorm"
	"ziyi.com/yiDa/model"
)

type userActivityDao struct{}

var UserActivityDao = new(userActivityDao)

func (d *userActivityDao) GetUserActivityByUidAndAid(engine *xorm.Engine, uid, aid int64) (int64, error) {
	i, err := engine.Where("uid = ? and aid = ?", uid, aid).Count()
	if err != nil {
		return 0, err
	}
	return i, nil
}

func (d *userActivityDao) GetUserActivityByAid(engine *xorm.Engine, uid int64) (userActivities []*model.UserActivity, err error) {
	err = engine.Where("uid = ?", uid).Find(&userActivities)
	return
}
