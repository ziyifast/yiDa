package main

import (
	"github.com/kataras/iris/v12"
	"ziyi.com/yiDa/controller"
	"ziyi.com/yiDa/pg"
)

func main() {
	app := iris.New()
	pg.InitDb()
	controller.InitControllers(app)
	app.Run(iris.Addr(":8083"))

	// engine := pg.Engine
	// user := &model.User{
	// 	Username:    "yiDa",
	// 	Password:    "fadsfasfs",
	// 	Wechat:      "fsafas",
	// 	Phone:       "41234124124",
	// 	City:        "xiamen",
	// 	Description: "lalala",
	// 	CreateTime:  time.Now(),
	// 	UpdateTime:  time.Now(),
	// }
	// dao.UserDao.Insert(engine, user)
}
