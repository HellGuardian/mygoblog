package controllers

import (
	"github.com/astaxie/beego"
	"mygoblog/models"
)

type TopicController struct {
	beego.Controller
}

func (this *TopicController) Get() {
	this.Data["IsLogin"] = checkAccount(this.Ctx)
	this.Data["IsTopic"] = true
	this.TplName = "topic.html"

	var err error
	this.Data["Topics"], err = models.GetAllTopic()
	if err != nil {
		beego.Error(err)
	}
}
