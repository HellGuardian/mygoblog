package controllers

import (
	"github.com/astaxie/beego"
	"strings"
	"mygoblog/models"
	"github.com/astaxie/beego/validation"
	"time"
)

type UserController struct {
	beego.Controller
}


func (this *UserController) Add() {
	this.TplName = "register.html"

	var num = make([]int, 0, 0)
	for i := 14; i <= 125; i++ {
		num = append(num, i)
	}
	this.Data["Number"] = num
	if this.Ctx.Request.Method == "POST" {
		errmsg := make(map[string]string)
		uid := this.Input().Get("uid")
		username := strings.TrimSpace(this.Input().Get("username"))
		password := strings.TrimSpace(this.Input().Get("password"))
		//password1 := strings.TrimSpace(this.Input().Get("password1"))
		sex := strings.TrimSpace(this.Input().Get("sex"))
		age := strings.TrimSpace(this.Input().Get("age"))
		const timeFormat = "2006-01-02"
		birthday, err := time.Parse(timeFormat, this.Input().Get("birthday"))
		if err != nil {
			beego.Error(err)
			return
		}
		email := strings.TrimSpace(this.Input().Get("email"))

		// 验证字段的有效性
		valid := validation.Validation{}
		if v := valid.Required(username, "username"); !v.Ok {
			errmsg["username"] = "请输入用户名。"
		} else if v := valid.MaxSize(username, 20, "username"); !v.Ok {
			errmsg["username"] = "用户名长度不能大于20个字符。"
		}

		if len(uid) == 0 {
			err = models.AddUser(username, password, sex, age, birthday, email)
		} else {
			err = models.EditUser(uid, username, password, sex, age, birthday, email)
		}
		if err != nil {
			beego.Error(err)
		}

		this.Redirect("/user/view", 302)
		return
	}
}

func (this *UserController) View() {
	this.TplName = "user_view.html"

	user, err := models.GetAllUsers()
	if err != nil {
		beego.Error(err)
	}

	this.Data["IsLogin"] = checkAccount(this.Ctx)
	this.Data["IsUser"] = true
	this.Data["User"] = user
	this.Data["IsEdit"] = false
}

// 删除用户
func (this *UserController) Delete() {
	if !checkAccount(this.Ctx) {
		this.Redirect("/login", 302)
		return
	}

	err := models.DeleteUser(this.Input().Get("uid"))
	if err != nil {
		beego.Error(err)
	}
	this. Redirect("/user/view", 302)
}

// 修改用户信息
func (this *UserController) Edit() {
	if !checkAccount(this.Ctx) {
		this.Redirect("/login", 302)
		return
	}
	this.TplName = "user_edit.html"
	uid := this.Ctx.Input.Param("0")
	user, err := models.GetUser(uid)
	if err != nil {
		beego.Error(err)
		this.Redirect("/user/view", 302)
		return
	}
	this.Data["User"] = user
	this.Data["Uid"] = uid

	if this.Ctx.Request.Method == "POST" {
		uid := this.Input().Get("uid")
		username := strings.TrimSpace(this.Input().Get("username"))
		password := strings.TrimSpace(this.Input().Get("password"))
		sex := strings.TrimSpace(this.Input().Get("sex"))
		age := strings.TrimSpace(this.Input().Get("age"))
		email := strings.TrimSpace(this.Input().Get("email"))
		const timeForm = "2016-01-01"
		birthday, err := time.Parse(timeForm, this.Input().Get("birthday"))
		if err != nil {
			beego.Error(err)
			return
		}

		err = models.EditUser(uid, username, password, sex, age, birthday, email)
		if err != nil {
			beego.Error(err)
			return
		}

		this.Redirect("/user/view", 302)
		return
	}
}