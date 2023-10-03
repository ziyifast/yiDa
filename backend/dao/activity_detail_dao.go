package dao

import (
	"xorm.io/xorm"
	"ziyi.com/yiDa/model"
)

type activityDetailDao struct{}

var ActivityDetailDao = new(activityDetailDao)

func (d *activityDetailDao) Insert(engine xorm.Interface, activityDetail *model.ActivityDetail) (int64, error) {
	return engine.Insert(activityDetail)
}

func (d *activityDetailDao) Get(engine xorm.Interface) (detailList []*model.ActivityDetail, err error) {
	err = engine.Find(&detailList)
	return
}
