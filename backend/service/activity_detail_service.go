package service

import (
	"encoding/json"

	"github.com/aobco/log"
	"ziyi.com/yiDa/dao"
	"ziyi.com/yiDa/dto"
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

func (s *activityDetailService) GetAll() ([]*dto.ActivityDetailDto, error) {
	detailList, err := dao.ActivityDetailDao.Get(pg.Engine)
	if err != nil {
		return nil, err
	}
	return transActivityDetailToDto(detailList)
}

func transActivityDetailToDto(detailList []*model.ActivityDetail) ([]*dto.ActivityDetailDto, error) {
	b, err2 := json.Marshal(detailList)
	if err2 != nil {
		return nil, err2
	}
	var detailDtoList []*dto.ActivityDetailDto
	err := json.Unmarshal(b, &detailDtoList)
	if err != nil {
		return nil, err
	}
	//封装报名人员列表
	userListMap, err := UserActivityService.GetUserListByAid(detailList)
	if err != nil {
		return nil, err
	}
	for i := 0; i < len(detailList); i++ {
		detailDtoList[i].UserList = userListMap[int64(detailDtoList[i].Id)]
	}
	return detailDtoList, nil
}

func (s *activityDetailService) GetActivityListByPage(startPage, pageSize int) ([]*dto.ActivityDetailDto, error) {
	detailList, err := dao.ActivityDetailDao.GetActivityListByPage(pg.Engine, startPage, pageSize)
	if err != nil {
		log.Errorf("%v", err)
		return nil, err
	}
	return transActivityDetailToDto(detailList)
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
