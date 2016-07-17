package models

import (
	"time"

	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/Unknwon/com"
	"strconv"
)

/* const(
	_DB_Name = "data/beelog.db"
	_MySql_DRIVER = "MySql"
)

var (
	dbuser string = "root"
	dbpasswd string = "linux"
    dbname string = "mygoblog"
	//dbhost string = "192.168.31.176"
	dbhost string = "192.168.191.3"
	dbport string = "3306"
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
	Category string
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
	//conn := dbuser + ":" + dbpasswd + "@tcp(" + dbhost + ":" + dbport + ")/" + dbname + "?charset=utf8&loc=Asia%2FShanghai"
	//orm.RegisterDataBase("default", "mysql", conn)
	//conn := dbuser + ":" + dbpasswd + "@tcp(" + dbhost + ":" + dbport + ")/" + dbname + "?charset=utf8&loc=Asia%2FShanghai"
	//orm.RegisterDataBase("default", "mysql", "mygoblog:linux@tcp(192.168.191.2:3306)/mygoblog?charset=utf8")
	orm.RegisterDataBase("default", "mysql", "mygoblog:linux@tcp(192.168.31.165:3306)/mygoblog?charset=utf8")
}

//添加分类
func AddCategory (name string) error {
	o := orm.NewOrm()

	cate := &Category{Title:name}

	qs := o.QueryTable("Category")
	err := qs.Filter("title", name).One(cate)
	if err == nil {
		return err
	}

	_, err = o.Insert(cate)
	if err != nil {
		return err
	}

	return nil
}

//获取所有的分类信息
func GetAllCategory() ([]*Category, error) {
	o := orm.NewOrm()

	cates := make([]*Category, 0)
	qs := o.QueryTable("Category")
	_, err := qs.All(&cates)
	return cates, err
}

//通过id删除分类
func DelCategory(id string) error {
	cid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return nil
	}

	o := orm.NewOrm()
	cate := &Category{Id: cid}
	_, err =o.Delete(cate)
	return err
}

// 添加文章的方法,直接写入数据库
func AddTopic(title, category, content string) error {
	o := orm.NewOrm()

	topic := &Topic{
		Title: title,
		Category: category,
		Content: content,
		Created: time.Now(),
		Updated: time.Now(),
	}

	_, err := o.Insert(topic)
	return err
}

//获取所有的文章信息
func GetAllTopics(isDesc bool) ([]*Topic, error) {
	o := orm.NewOrm()

	topics := make([]*Topic, 0)
	qs := o.QueryTable("Topic")

	var err error
	if isDesc {
		_, err = qs.OrderBy("-created").All(&topics)
	} else {
		_, err = qs.All(&topics)
	}
	return topics, err
}

// 浏览文章
func GetTopic(tid string) (*Topic, error) {
	tidNum, err := strconv.ParseInt(tid, 10, 64)
	if err != nil {
		return nil, err
	}
	o := orm.NewOrm()
	topic := new(Topic)

	qs := o.QueryTable("topic")
	err = qs.Filter("id", tidNum).One(topic)
	if err != nil {
		return nil, err
	}
	topic.Views++
	_, err = o.Update(topic)
	return topic, err
}

// 修改文章
func ModifyTopic(tid, title, category, content string) error {
	tidNum, err := strconv.ParseInt(tid, 10, 64)
	if err != nil {
		return err
	}
	o := orm.NewOrm()
	topic := &Topic{Id: tidNum}
	if o.Read(topic) == nil {
		topic.Title = title
		topic.Category = category
		topic.Content = content
		topic.Updated = time.Now()
		o.Update(topic)
	}
	return err
}

// 删除文章的方法
func DeleteTopic(tid string) error {
	tidNum, err := strconv.ParseInt(tid, 10, 64)
	if err != nil {
		return err
	}

	o := orm.NewOrm()
	topic := &Topic{Id: tidNum}
	_, err = o.Delete(topic)
	return err
}