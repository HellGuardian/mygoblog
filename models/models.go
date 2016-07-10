package models

import (
	"time"

	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/Unknwon/com"
)

/* const(
	_DB_Name = "data/beelog.db"
	_MySql_DRIVER = "MySql"
)*/

var (
	dbuser string = "root"
	dbpasswd string = "linux"
    dbname string = "mygoblog"
	dbhost string = "192.168.31.176"
	dbport string = "3306"
)

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
	orm.RegisterDriver("mysql", orm.DRMySQL)

	// 注册默认数据库,数据库名为mygoblog,密码为linux
	//conn := dbuser + ":" + dbpasswd + "@/" + dbname + "?charset=utf8"
	//orm.RegisterDataBase("default", "mysql", conn)
	conn := dbuser + ":" + dbpasswd + "@tcp(" + dbhost + ":" + dbport + ")/" + dbname + "?charset=utf8&loc=Asia%2FShanghai"
	orm.RegisterDataBase("default", "mysql", conn)
}
