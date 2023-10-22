package dao

import (
	"github.com/aobco/log"
	"xorm.io/xorm"
	"ziyi.com/yiDa/model"
)

type activityDetailDao struct{}

var ActivityDetailDao = new(activityDetailDao)

func (d *activityDetailDao) Insert(engine xorm.Interface, activityDetail *model.ActivityDetail) (int64, error) {
	return engine.Omit("id").Insert(activityDetail)
}

func (d *activityDetailDao) Get(engine xorm.Interface) (detailList []*model.ActivityDetail, err error) {
	err = engine.OrderBy("id").Find(&detailList)
	if err != nil {
		log.Errorf("%v", err)
		return
	}
	return
}

func (d *activityDetailDao) GetActivityListByPage(engine xorm.Interface, startPage, pageSize int) (detailList []*model.ActivityDetail, err error) {
	err = engine.Limit(pageSize, (startPage-1)*pageSize).OrderBy("id").Find(&detailList)
	return
}

func (d *activityDetailDao) UpdateActivityDetailById(engine xorm.Interface, activityDetail *model.ActivityDetail) (int64, error) {
	return engine.Where("id=?", activityDetail.Id).Update(activityDetail)
}

func (d *activityDetailDao) GetActivityDetailById(engine xorm.Interface, id int64) (*model.ActivityDetail, error) {
	detail := new(model.ActivityDetail)
	_, err := engine.Where("id=?", id).Get(detail)
	if err != nil {
		log.Errorf("%v", err)
		return nil, err
	}
	return detail, nil
}

func (d *activityDetailDao) GetActivitiesByIds(engine xorm.Interface, ids []int64) (detailList []*model.ActivityDetail, err error) {
	err = engine.In("id", ids).Find(&detailList)
	if err != nil {
		log.Errorf("%v", err)
		return
	}
	return
}
