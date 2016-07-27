package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

type LoginController struct  {
	beego.Controller
}

func (this *LoginController) Get() {
	//sess, _ := beego.GlobalSessions.SessionStart()
	isExit := this.Input().Get("exit") == "true"
	if isExit {
		//this.DelSession("username")
		//this.DelSession("pwd")
		this.Ctx.SetCookie("username", "", -1, "/")
		this.Ctx.SetCookie("pwd", "", -1, "/")
		this.Redirect("/login.html", 301)
		return
	}
	this.TplName = "login.html"
}

func (this *LoginController) Post() {
	username := this.Input().Get("username")
	pwd := this.Input().Get("pwd")
	autoLogin := this.Input().Get("autoLogin") == "on"
	if beego.AppConfig.String("username") == username && beego.AppConfig.String("pwd") == pwd {
		maxAge := 0
		if autoLogin {
			maxAge = 1 << 32 - 1
		}

		this.Ctx.SetCookie("username", username, maxAge, "/")
		this.Ctx.SetCookie("pwd", pwd, maxAge, "/")
	} else {
		this.Data["errmsg"] = "用户名或密码错误,请从新输入!"
		this.Redirect("/login", 301)
		return
	}
	this.Redirect("/", 301)
	return
}


func checkAccount(ctx *context.Context) bool {
	ck, err := ctx.Request.Cookie("username")
	if err != nil {
		return false
	}
	username := ck.Value

	ck, err = ctx.Request.Cookie("pwd")
	if err != nil {
		return false
	}
	pwd := ck.Value

	return beego.AppConfig.String("username") == username && beego.AppConfig.String("pwd") == pwd
}
