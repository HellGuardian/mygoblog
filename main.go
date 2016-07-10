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

	// 运行时
	beego.Router("/", &controllers.MainController{})
	beego.Run()
}
