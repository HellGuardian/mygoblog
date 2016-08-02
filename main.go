package main

import (
	_ "mygoblog/controllers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_"github.com/go-sql-driver/mysql"
	"mygoblog/models"
	"mygoblog/controllers"
)

// 引入数据模型
func init() {
	// 注册数据库
	models.RegisterDB()
}

func main() {
	// 开启 ORM 调式模式
	orm.Debug = true

	// 自动建表
	orm.RunSyncdb("default", false, true)

	// 注册路由
	beego.Router("/", &controllers.MainController{})
	beego.Router("/login", &controllers.LoginController{})
	beego.Router("/category", &controllers.CategoryController{})
	beego.Router("/topic", &controllers.TopicController{})
	beego.AutoRouter(&controllers.TopicController{})
	beego.Router("/reply", &controllers.ReplyController{})
	beego.Router("/reply/add", &controllers.ReplyController{}, "post:Add")
	beego.Router("/reply/delete", &controllers.ReplyController{}, "get:Delete")
	beego.Router("/register", &controllers.UserController{}, "*:Add")
	beego.Router("/user/view", &controllers.UserController{}, "*:View")
	beego.Router("/user/delete",&controllers.UserController{}, "get:Delete")
	beego.Router("/user/edit", &controllers.UserController{}, "get:Edit")
	beego.Router("/user/edit", &controllers.UserController{}, "post:Add")
	beego.AutoRouter(&controllers.UserController{})
	beego.Run()
}
