package pg

import (
	"github.com/aobco/log"
	_ "github.com/lib/pq"
	"xorm.io/xorm"
)

var (
	dbUrl = "host=localhost port=5432 user=postgres password=postgres dbname=yida sslmode=disable"
)

var Engine *xorm.Engine

func InitDb() {
	//初始化数据库配置
	engine, err := xorm.NewEngine("postgres", dbUrl)
	if err != nil {
		log.Errorf("%v", err)
	}
	engine.ShowSQL(true)
	// fmt.Println("db=", engine)
	Engine = engine
}
