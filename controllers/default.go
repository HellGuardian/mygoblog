package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct  {
	beego.Controller
}

func (this *MainController) Get() {
	// 定义首页模板
	this.TplName = "home.html"
}