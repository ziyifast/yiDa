package controller

import (
	"net/http"

	"github.com/aobco/log"
	"github.com/kataras/iris/v12/mvc"
	"ziyi.com/yiDa/service"
)

type UserActivityController struct {
	BaseController
}

func (u *UserActivityController) BeforeActivation(b mvc.BeforeActivation) {
	b.Handle(http.MethodPost, "/delete/{uid:int}/{aid:int}", "DelUserActivity")
	b.Handle(http.MethodPost, "/add/{uid:int}/{aid:int}", "AddUserActivity")
	b.Handle(http.MethodGet, "/smoke", "Smoke")
}

// func (c *UserActivityController) DeleteUserActivityByUidAndAid() mvc.Result {
// 	uid, err := c.ParamInt("uid")
// 	if err != nil {
// 		return c.BadRequest()
// 	}
// 	aid, err := c.ParamInt("aid")
// 	if err != nil {
// 		return c.BadRequest()
// 	}
// 	err2 := service.UserActivityService.DeleteUserActivityByUidAndAid(int64(uid), int64(aid))
// 	if err2 != nil {
// 		log.Errorf("%v", err2)
// 		return c.Failed(err2.Error())
// 	}
// 	return c.Ok(nil)

// }

func (c *UserActivityController) AddUserActivity() mvc.Result {
	uid, err := c.ParamInt("uid")
	if err != nil {
		return c.BadRequest()
	}
	aid, err := c.ParamInt("aid")
	if err != nil {
		return c.BadRequest()
	}
	//查询活动状态是否已经暂停或关闭（status是否为0）
	ad, err2 := service.ActivityDetailService.GetActivityDetailById(int64(aid))
	if err2 != nil {
		log.Errorf("%v", err2)
		return c.SystemInternalErrorWithMsg(err2.Error())
	}
	if ad.Status == 0 {
		return c.SystemInternalErrorWithMsg("活动已暂停或已关闭")
	}
	//查询是否已经加入了该活动
	flag, err2 := service.UserActivityService.GetUserActivityByUidAndAid(int64(uid), int64(aid))
	if err2 != nil {
		log.Errorf("%v", err2)
		return c.Failed(err2.Error())
	}
	if flag {
		return c.SystemInternalErrorWithMsg("你已经报名了该活动，不可重复报名")
	}
	err = service.UserActivityService.AddUserActivity(int64(uid), int64(aid))
	if err != nil {
		log.Errorf("%v", err)
		return c.Failed(err.Error())
	}
	return c.Ok(nil)
}

func (c *UserActivityController) DelUserActivity() mvc.Result {
	uid, err := c.ParamInt("uid")
	if err != nil {
		return c.BadRequest()
	}
	aid, err := c.ParamInt("aid")
	if err != nil {
		return c.BadRequest()
	}
	err = service.UserActivityService.DelUserActivity(int64(uid), int64(aid))
	if err != nil {
		log.Errorf("%v", err)
		return c.Failed(err.Error())
	}
	return c.Ok(nil)
}

func (c *UserActivityController) Smoke() mvc.Result {
	return c.Ok("smoke...")
}
