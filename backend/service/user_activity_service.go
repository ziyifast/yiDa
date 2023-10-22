package service

import (
	"errors"
	"strconv"
	"strings"

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
	uidStr := strconv.Itoa(int(uid))
	activityDetail.UserList = activityDetail.UserList + "," + uidStr
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
	arr := strings.Split(activityDetail.UserList, ",")
	list := make([]string, 0)
	for _, v := range arr {
		if v != strconv.Itoa(int(uid)) {
			list = append(list, v)
		}
	}
	activityDetail.UserList = strings.Join(list, ",")
	log.Infof("activityDetail:%v", activityDetail)
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
