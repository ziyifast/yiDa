package main

import (
	"net/http"

	"github.com/kataras/iris/v12"
	"github.com/rs/cors"
	"ziyi.com/yiDa/controller"
	"ziyi.com/yiDa/pg"
)

func main() {
	app := iris.New()
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:8080", "http://localhost:8083"}, //填写实际的允许来源
		AllowedMethods:   []string{http.MethodOptions, http.MethodPost, http.MethodGet},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: true,
		Debug:            true,
	})
	app.WrapRouter(c.ServeHTTP)

	pg.InitDb()
	controller.InitControllers(app)
	app.Run(iris.Addr(":8083"))
}
