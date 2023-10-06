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
