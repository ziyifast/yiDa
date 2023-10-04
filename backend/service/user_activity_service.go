package service

import (
	"ziyi.com/yiDa/dao"
	"ziyi.com/yiDa/model"
	"ziyi.com/yiDa/pg"
)

type userActivityService struct{}

var UserActivityService = new(userActivityService)

func (s *userActivityService) GetUserActivitiesByAid(aId int) (userActivity []*model.UserActivity, err error) {
	userActivity, err = dao.UserActivityDao.GetUserActivityByAid(pg.Engine, int64(aId))
	return
}
