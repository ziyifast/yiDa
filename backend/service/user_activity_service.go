package service

import (
	"errors"

	"github.com/aobco/log"
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

func (s *userActivityService) AddUserActivity(uid, aid int64) (err error) {
	defer func() {
		if e := recover(); e != nil {
			log.Errorf("%v", e)
			err = errors.New("transaction err")
		}
	}()
	session := pg.Engine.NewSession()
	session.Begin()
	defer session.Close()
	userActivity := &model.UserActivity{
		Uid: uid,
		Aid: aid,
	}
	err = dao.UserActivityDao.InsertUserActivity(session, userActivity)
	if err != nil {
		log.Errorf("%v", err)
		session.Rollback()
		return err
	}
	//更新activityDetail
	activityDetail, err := dao.ActivityDetailDao.GetActivityDetailById(session, aid)
	log.Infof("activityDetail:%v", activityDetail)
	if err != nil {
		log.Errorf("%v", err)
		session.Rollback()
		return err
	}
	_, err = dao.ActivityDetailDao.UpdateActivityDetailById(session, activityDetail)
	if err != nil {
		log.Errorf("%v", err)
		session.Rollback()
		return err
	}
	session.Commit()
	return nil
}

func (s *userActivityService) DelUserActivity(uid, aid int64) (err error) {
	defer func() {
		if e := recover(); e != nil {
			log.Errorf("%v", e)
			err = errors.New("transaction err")
		}
	}()
	session := pg.Engine.NewSession()
	session.Begin()
	defer session.Close()
	err = dao.UserActivityDao.DeleteByUidAndAid(session, uid, aid)
	if err != nil {
		log.Errorf("%v", err)
		session.Rollback()
		return err
	}
	//更新activityDetail
	activityDetail, err := dao.ActivityDetailDao.GetActivityDetailById(session, aid)
	if err != nil {
		log.Errorf("%v", err)
		return err
	}
	_, err = dao.ActivityDetailDao.UpdateActivityDetailById(session, activityDetail)
	if err != nil {
		log.Errorf("%v", err)
		session.Rollback()
		return err
	}
	session.Commit()
	return nil
}

func (s *userActivityService) GetUserActivityByUidAndAid(uid, aid int64) (bool, error) {
	if uid < 0 || aid < 0 {
		return false, errors.New("invalid param")
	}
	flag, err := dao.UserActivityDao.GetUserActivityDetailByUidAndAid(*pg.Engine, uid, aid)
	if err != nil {
		return false, err
	}
	return flag, nil
}

func (s *userActivityService) GetUserListByAid(detailList []*model.ActivityDetail) (map[int64]string, error) {
	res := make(map[int64]string)
	//获取活动列表与参与人员的关联关系
	for _, v := range detailList {
		ids, err := dao.UserActivityDao.GetUidsByAid(pg.Engine, int64(v.Id))
		if err != nil {
			log.Errorf("根据活动id获取报名人员列表失败%v", err)
			return nil, err
		}
		var userListStr string
		for _, uid := range ids {
			user, err := dao.UserDao.GetUserById(pg.Engine, uid)
			if err != nil {
				log.Errorf("根据uid获取用户信息失败%v", err)
				return nil, err
			}
			userListStr += user.Username + ","
		}
		res[int64(v.Id)] = userListStr
	}
	return res, nil
}
