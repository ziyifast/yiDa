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
