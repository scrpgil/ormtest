package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/mattn/go-sqlite3"
	"time"
)

type User struct {
	Id           int
	Name         string     `orm:"column(user_name);size(30)"`
	Email        string     `orm:"null"`
	Subscription int        `orm:"default(1)"`
	Comment      []*Comment `orm:"reverse(many)"`
	Created      time.Time  `orm:"auto_now_add;type(datetime)"`
	Updated      time.Time  `orm:"auto_now;type(datetime)"`
}

type Comment struct {
	Id    int
	User  *User `orm:"rel(fk)"`
	Title string
	Msg   string
}

func (u *User) TableName() string {
	return "auth_user"
}

func init() {
	// Need to register model in init
	orm.RegisterDriver("sqlite", orm.DRSqlite)
	err := orm.RegisterDataBase("default", "sqlite3", "file:data.db")
	if err != nil {
		panic(err)
	}

	orm.RegisterModel(new(User))
	orm.RegisterModel(new(Comment))
	// Drop table and re-create.
	force := false
	// Print log.
	verbose := false
	err = orm.RunSyncdb("default", force, verbose)
	if err != nil {
		beego.Error(err)
	}
}

func AddComment(title string, user *User) error {
	o := orm.NewOrm()
	o.Using("default")
	c := Comment{Title: title, User: user}
	_, err := o.Insert(&c)
	if err != nil {
		beego.Error("err:", err)
		return err
	}

	return nil

}
func AddUser() error {
	o := orm.NewOrm()
	o.Using("default")
	user := User{Id: 1, Email: "test@test.com"}
	_, err := o.Insert(&user)
	if err != nil {
		beego.Error("err:", err)
		return err
	}

	return nil

}
func GetUser() (User, error) {
	o := orm.NewOrm()
	o.Using("default")
	user := User{Id: 1}
	//err := o.RelatedSel(&user)
	err := o.QueryTable("user").Filter("id", 1).RelatedSel().One(&user)
	if err != nil {
		beego.Error("err:", err)
		return user, err
	}

	return user, nil
}
