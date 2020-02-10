package main

import (
	"fg-admin/config"
	"fg-admin/controller"
	"fg-admin/middleware"
	models "fg-admin/model"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/sessions"

	"github.com/kataras/iris/v12/middleware/logger"
	"github.com/kataras/iris/v12/middleware/recover"
	"github.com/kataras/iris/v12/mvc"
)

func main() {
	app := iris.New()

	// 同步数据表
	models.DB.AutoMigrate(
		&models.User{},
		&models.Role{},
		&models.Permission{},
		&models.OperationLog{},
		&models.Season{},
	)

	iris.RegisterOnInterrupt(func() {
		_ = models.DB.Close()
	})

	// 注册静态文件目录
	app.RegisterView(iris.HTML("./views", ".html").Layout("public/layout.html"))
	app.HandleDir("/static", "./static", iris.DirOptions{
		ShowList: true,
	})

	app.Logger().SetLevel("debug")
	// app.Logger().AddOutput()

	// 中间件
	app.Use(recover.New())
	// 敏感操作日志，记入数据库
	app.Use(middleware.NewOperateLogger())
	app.Use(logger.New())
	app.Use(middleware.NewAuth())


	mvc.Configure(app.Party("/"), basicMVC)
	app.Run(iris.Addr(config.Conf.Get("app.addr").(string)), iris.WithoutServerError(iris.ErrServerClosed))
}

func basicMVC(app *mvc.Application) {
	app.Register(
		sessions.New(sessions.Config{}).Start,
	)

	app.Party("/home").Handle(new(controller.HomeController))
	app.Party("/user").Handle(new(controller.UserController))
	app.Party("/server").Handle(new(controller.ServerController))
	app.Party("/role").Handle(new(controller.RoleController))
	app.Party("/auth").Handle(new(controller.AuthController))
	app.Party("/season").Handle(new(controller.SeasonController))

}
