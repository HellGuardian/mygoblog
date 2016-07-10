package models

import (
	"time"

	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	//"github.com/Unknwon/com"
)

/* const(
	_DB_Name = "data/beelog.db"
	_MySql_DRIVER = "MySql"
)*/

type Category struct {
	Id	int64
	Title	string
	Created	time.Time	`orm:"index"`
	Views	int64	`orm:"index"`
	TopicCount	int64
	TopicTime	time.Time	`orm:"index"`
	TopicLastUserId	int64
}

type Topic struct {
	Id	int64
	Uid	int64
	Title	string
	Content	string	`orm:"size(5000)"`
	Attachment	string
	Created	time.Time	`orm:"index"`
	Updated	time.Time	`orm:"index"`
	Views	int64	`orm:"index"`
	Author	string
	ReplyTime	time.Time	`orm:"index"`
	ReplyCount	int64
	ReplyLastUserId	int64
}

func RegisterDB() {
	// 注册model
	orm.RegisterModel(new(Category), new(Topic))

	// 注册驱动
	orm.RegisterDriver("mysql", orm.DR_MySQL)

	// 注册默认数据库  密码为空的格式
	orm.RegisterDataBase("default", "mysql", "root:@/app?charset=utf8")
}
