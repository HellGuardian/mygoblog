package models

import (
	"github.com/astaxie/beego/orm"
	"time"
	"strconv"
)

func AddUser(username, password, sex, age string, birthday time.Time, email string,) error {
	o := orm.NewOrm()
	s_age, err := strconv.ParseInt(age, 10, 64)
	if err != nil {
		return err
	}

	user := &User {
		UserName: username,
		PassWord: password,
		Sex: sex,
		Age: s_age,
		Birthday: birthday,
		Email: email,
		RegisterTime: time.Now(),
	}
	_, err = o.Insert(user)
	return err
}

func GetAllUsers() ([]*User, error) {
	o := orm.NewOrm()

	user := make([]*User, 0)
	qs := o.QueryTable("User")
	_, err := qs.All(&user)
	return user, err
}

// 通过用户id删了用户
func DeleteUser(id string) error {
	uid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return nil
	}

	o := orm.NewOrm()
	user := &User{Id:uid}
	_, err = o.Delete(user)
	return nil
}

// 修改用户信息
func EditUser(id, username, password, sex, age string, birthday time.Time, email string) error {
	uid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return nil
	}
	o := orm.NewOrm()
	user := &User{Id: uid}
	if o.Read(user) == nil {
		user.UserName = username
		user.PassWord = password
		user.Sex = sex
		user.Age, err = strconv.ParseInt(age, 10, 64)
		if err != nil {
			return nil
		}
		user.Birthday = birthday
		user.Email = email
		o.Update(user)
	}
	return nil
}

func GetUser(uid string) (*User, error) {
	uidNum, err := strconv.ParseInt(uid, 10, 64)
	if err != nil {
		return nil, err
	}
	o := orm.NewOrm()
	user := new(User)

	qs := o.QueryTable(user)
	err = qs.Filter("id", uidNum).One(user)
	if err != nil {
		return nil, err
	}
	return user, err
}