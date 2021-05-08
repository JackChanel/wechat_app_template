package main

import (
	"os"
	"strconv"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/jinzhu/gorm"
	"gopkg.in/kataras/iris.v6"
	"gopkg.in/kataras/iris.v6/adaptors/httprouter"
	"gopkg.in/kataras/iris.v6/adaptors/sessions"
	"wemall/go/config"
	"wemall/go/model"
	"wemall/go/route"
)

func main() {
	db, err := gorm.Open(config.DBConfig.Dialect, config.DBConfig.URL)
	if err != nil {
		os.Exit(1)
	}

	if config.DBConfig.SQLLog {
		db.LogMode(true)
	}

	db.DB().SetMaxIdleConns(config.DBConfig.MaxIdleConns);
	db.DB().SetMaxOpenConns(config.DBConfig.MaxOpenConns)

	model.DB = db;

	app := iris.New(iris.Configuration{
        Gzip    : true, 
        Charset : "UTF-8",
	})

	if config.ServerConfig.Debug {
		app.Adapt(iris.DevLogger())
	}

	app.Adapt(httprouter.New())

	app.Adapt(sessions.New(sessions.Config{
		Cookie: config.ServerConfig.SessionID,
	}))

	route.Route(app)

	app.OnError(iris.StatusNotFound, func(ctx *iris.Context) {
		ctx.JSON(iris.StatusOK, iris.Map{
			"errNo" : model.ErrorCode.NotFound,
			"msg"   : "Not Found",
			"data"  : iris.Map{},
		})

	})

	app.OnError(500, func(ctx *iris.Context) {
		ctx.JSON(iris.StatusOK, iris.Map{
			"errNo" : model.ErrorCode.ERROR,
			"msg"   : "error",
			"data"  : iris.Map{},
		})
	})

	app.Listen(":" + strconv.Itoa(config.ServerConfig.Port))
}
