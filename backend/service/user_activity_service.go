package service

import (
	"ziyi.com/yiDa/dao"
	"ziyi.com/yiDa/model"
	"ziyi.com/yiDa/pg"
)

type userActivityService struct{}

var UserActivityService = new(userActivityService)

func (s *userActivityService) GetUserActivitiesByUid(uId int) (userActivity []*model.UserActivity, err error) {
	userActivity, err = dao.UserActivityDao.GetUserActivityByUid(pg.Engine, int64(uId))
	return
}

func (s *userActivityService) DeleteUserActivityByUidAndAid(uid, aid int64) (err error) {
	// 删除用户活动
	err = dao.UserActivityDao.DeleteByUidAndAid(pg.Engine, uid, aid)
	return
}
