package controller

import (
	"net/http"

	"github.com/aobco/log"
	"github.com/kataras/iris/v12/mvc"
	"ziyi.com/yiDa/model"
	"ziyi.com/yiDa/service"
)

type ActivityDetailController struct {
	BaseController
}

func (c *ActivityDetailController) BeforeActivation(b mvc.BeforeActivation) {
	b.Handle(http.MethodGet, "/list", "GetAllActivityDetail")
	b.Handle(http.MethodPost, "/add", "AddActivityDetail")
	b.Handle(http.MethodPost, "/update", "UpdateActivityDetail")
	b.Handle(http.MethodGet, "/get/{uid:int64}", "GetUserActivitiesByUid")
}

func (c *ActivityDetailController) AddActivityDetail() mvc.Result {
	activityDetail := &model.ActivityDetail{}
	err := c.ReadJSON(activityDetail)
	if err != nil {
		log.Errorf("%v", err)
		return c.Failed(err.Error())
	}

	err = service.ActivityDetailService.Insert(activityDetail)
	if err != nil {
		log.Errorf("%v", err)
		return c.Failed(err.Error())
	}
	return c.Ok("add success...")
}

func (c *ActivityDetailController) GetAllActivityDetail() mvc.Result {
	list, err := service.ActivityDetailService.GetAll()
	if err != nil {
		log.Errorf("%v", err)
		return c.Failed(err.Error())
	}
	return c.Ok(list)
}

func (c *ActivityDetailController) UpdateActivityDetail() mvc.Result {
	activityDetail := &model.ActivityDetail{}
	if err := c.ReadJSON(activityDetail); err != nil {
		log.Errorf("%v", err)
		return c.Failed(err.Error())
	}
	log.Infof("%v", activityDetail)
	err := service.ActivityDetailService.UpdateActivityDetailById(activityDetail)
	if err != nil {
		log.Errorf("%v", err)
		return c.Failed(err.Error())
	}
	return c.Ok("update success...")
}

/*
根据用户id获取用户参加的活动信息
*/
func (c *ActivityDetailController) GetUserActivitiesByUid() mvc.Result {
	uid, err := c.ParamInt64("uid")
	if err != nil {
		log.Errorf("%v", err)
		return c.Failed(err.Error())
	}
	//获取活动ids
	userActivityIds, err2 := service.UserActivityService.GetUserActivitiesByUid(int(uid))
	if err2 != nil {
		log.Errorf("%v", err2)
		return c.Failed(err2.Error())
	}
	aIds := make([]int64, 0)
	for _, v := range userActivityIds {
		aIds = append(aIds, v.Aid)
	}

	//根据活动ids，查询出活动详情信息
	activities, err3 := service.ActivityDetailService.GetActivitiesByIds(aIds)
	if err3 != nil {
		log.Errorf("%v", err3)
		return c.Failed(err3.Error())
	}
	return c.Ok(activities)
}
