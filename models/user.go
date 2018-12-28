package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type User struct {
	Id         int64
	Username   string    `orm:"unique;size(15)"`
	Password   string    `orm:"size(32)"`
	Email      string    `orm:"size(50)"`
	Lastlogin  time.Time `orm:"auto_now_add;type(datetime)"`
	Logincount int64
	Lastip     string `orm:"size(32)"`
	Authkey    string `orm:"size(10)"`
	Active     int8
}

func (U *User) TableName()  string{
	return TableName("user")
}

func (U *User) Insert() error{
	if _,err := orm.NewOrm().Insert(U); err != nil{
		return err
	}
	return nil
}

func (U *User) Read(fields ...string) error {
	if err := orm.NewOrm().Read(U,fields...); err != nil {
		return err
	}
	return nil
}

func (U *User) Update(fields ...string) error  {
	if _, err := orm.NewOrm().Update(U,fields...); err != nil{
		return err
	}
	return nil
}
func (U *User) Delete() error{
	if _,err := orm.NewOrm().Delete(U); err != nil {
		return err
	}
	return nil
}
func (U *User) Query() orm.QuerySeter {
	return orm.NewOrm().QueryTable(U)
}
