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