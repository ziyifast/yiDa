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
	b.Handle(http.MethodPost, "/login/{username:string}/{password}", "Login")
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
	err := service.UserService.AddUser(user)
	if err != nil {
		return c.Failed(err.Error())
	}
	return c.Ok("")
}

func (c *UserController) Login() mvc.Result {
	log.Infof("params=%v\n", c.Ctx.Params())
	name := c.Param("username")
	pwd := c.Param("password")
	log.Infof("username=%v pwd=%v\n", name, pwd)
	u, err := service.UserService.GetUserByName(name)
	if err != nil {
		log.Errorf("%v", err)
		return c.Failed(err.Error())
	}
	if u == nil {
		return c.Failed("用户不存在")
	}

	if u.Password != pwd {
		return c.Failed("密码错误")
	}
	return c.Ok(u)
}

func (c *UserController) UpdateUser() mvc.Result {
	user := new(model.User)
	if err := c.ReadJSON(user); err != nil {
		return c.Failed(err.Error())
	}
	service.UserService.UpdateUser(user)
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
