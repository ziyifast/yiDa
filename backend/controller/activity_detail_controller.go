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
