package controller

import (
	"net/http"

	"github.com/aobco/log"
	"github.com/kataras/iris/v12/mvc"
	"ziyi.com/yiDa/model"
	"ziyi.com/yiDa/service"
)

type UserController struct {
	BaseController
}

func (c *UserController) BeforeActivation(b mvc.BeforeActivation) {
	b.Handle(http.MethodPost, "/add", "AddUser")
	b.Handle(http.MethodGet, "/smoke", "Smoke")
}

/**
新增用户
*/
func (c *UserController) AddUser() mvc.Result {
	user := new(model.User)
	if err := c.ReadJSON(user); err != nil {
		return c.Failed(err.Error())
	}
	// if err := c.Ctx.ReadJSON(user); err != nil {
	// 	log.Errorf("add user error:%v", err.Error())
	// 	return mvc.Response{
	// 		Code: http.StatusOK,
	// 		Object: map[string]interface{}{
	// 			"code":    0,
	// 			"message": fmt.Sprintf("添加用户失败, msg:%v", err.Error()),
	// 		},
	// 	}
	// }
	err := service.UserService.AddUser(user)
	if err != nil {
		return c.Failed(err.Error())
		// return mvc.Response{
		// 	Code: http.StatusOK,
		// 	Object: map[string]interface{}{
		// 		"code":    0,
		// 		"message": fmt.Sprintf("添加用户失败, msg:%v", err.Error()),
		// 	},
		// }
	}
	return c.Ok("add user success")
	// return mvc.Response{
	// 	Code: http.StatusOK,
	// 	Object: map[string]interface{}{
	// 		"code":    1,
	// 		"message": "添加用户成功",
	// 	},
	// }
}

func (c *UserController) UpdateUser() mvc.Result {
	user := new(model.User)
	if err := c.Ctx.ReadJSON(user); err != nil {
		return c.Failed(err.Error())
	}
	return c.Ok("")
}

func (c *UserController) Smoke() mvc.Result {
	log.Infof("smoke ctx:%v\n", c.Ctx)
	return mvc.Response{
		Code: http.StatusOK,
		Object: map[string]interface{}{
			"code":    1,
			"message": "ok",
		},
	}
}
