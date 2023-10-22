package service

import (
	"github.com/aobco/log"
	"ziyi.com/yiDa/dao"
	"ziyi.com/yiDa/model"
	"ziyi.com/yiDa/pg"
)

type activityDetailService struct{}

var ActivityDetailService = activityDetailService{}

func (s *activityDetailService) Insert(detail *model.ActivityDetail) error {
	_, err := dao.ActivityDetailDao.Insert(pg.Engine, detail)
	if err != nil {
		return err
	}
	log.Infof("插入detail success")
	return nil
}

func (s *activityDetailService) GetAll() ([]*model.ActivityDetail, error) {
	detailList, err := dao.ActivityDetailDao.Get(pg.Engine)
	if err != nil {
		return nil, err
	}
	return detailList, nil
}

func (s *activityDetailService) GetActivityListByPage(startPage, pageSize int) ([]*model.ActivityDetail, error) {
	detailList, err := dao.ActivityDetailDao.GetActivityListByPage(pg.Engine, startPage, pageSize)
	if err != nil {
		log.Errorf("%v", err)
		return nil, err
	}
	return detailList, nil
}

func (s *activityDetailService) UpdateActivityDetailById(detail *model.ActivityDetail) error {
	// 更新活动详情表
	_, err := dao.ActivityDetailDao.UpdateActivityDetailById(pg.Engine, detail)
	if err != nil {
		return err
	}
	return nil
}

func (s *activityDetailService) GetActivitiesByIds(ids []int64) (activities []*model.ActivityDetail, err error) {
	activities, err = dao.ActivityDetailDao.GetActivitiesByIds(pg.Engine, ids)
	if err != nil {
		return nil, err
	}
	return activities, nil
}

func (s *activityDetailService) GetActivityDetailById(id int64) (detail *model.ActivityDetail, err error) {
	detail, err = dao.ActivityDetailDao.GetActivityDetailById(pg.Engine, id)
	return
}
