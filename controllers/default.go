package controllers

import (
	"github.com/astaxie/beego"
	"mygoblog/models"
)

type MainController struct  {
	beego.Controller
}

func (this *MainController) Get() {
	// 定义首页模板
	this.Data["IsHome}"] = true
	this.TplName = "home.html"

	this.Data["IsLogin"] = checkAccount(this.Ctx)

	topics, err := models.GetAllTopics(this.Input().Get("cate"), true)
	if err != nil {
		beego.Error(err.Error())
	} else {
		this.Data["Topics"] = topics
	}

	category, err := models.GetAllCategory()
	if err != nil {
		beego.Error(err)
	}

	this.Data["Categories"] = category
}