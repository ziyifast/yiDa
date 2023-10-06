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
	err = engine.OrderBy("id").Find(&detailList)
	return
}

func (d *activityDetailDao) UpdateActivityDetailById(engine xorm.Interface, activityDetail *model.ActivityDetail) (int64, error) {
	return engine.Where("id=?", activityDetail.Id).Update(activityDetail)
}

func (d *activityDetailDao) GetActivitiesByIds(engine xorm.Interface, ids []int64) (detailList []*model.ActivityDetail, err error) {
	err = engine.In("id", ids).Find(&detailList)
	return
}
