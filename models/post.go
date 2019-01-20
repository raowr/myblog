package models

import (
	"github.com/astaxie/beego/orm"
	"strings"
	"time"
)

type Post struct {
	Id       int64
	Userid   int64  `orm:"index"`
	Author   string `orm:"size(15)"`
	Title    string `orm:"size(100)"`
	Color    string `orm:"size(7)"`
	Urlname  string `orm:"size(100);index"`
	Urltype  int8
	Content  string    `orm:"type(text)"`
	Brief    string    `orm:"type(text)"`
	Tags     string    `orm:"size(100)"`
	Posttime time.Time `orm:"type(datetime);index"`
	Views    int64
	Status   int8
	Updated  time.Time `orm:"type(datetime)"`
	Istop    int8
	Cover    string `orm:"size(70)"`
}

/** 返回表名 */
func (m *Post) TableName() string {
	return TableName("post")
}

/** 数据插入 */
func (m *Post) Insert() error {
	if _, err := orm.NewOrm().Insert(m); err != nil {
		return err
	}
	return nil
}

/** 读数据 */
func (m *Post) Read(fields ...string) error {
	if err := orm.NewOrm().Read(m, fields...); err != nil {
		return err
	}
	return nil
}

/** 更新数据 */
func (m *Post) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(m, fields...); err != nil {
		return err
	}
	return nil
}

/** 删除数据 */

func (m *Post) Delete() error {
	if m.Tags != "" {
		o := orm.NewOrm()
		oldtags := strings.Split(strings.Trim(m.Tags, ","), ",")
		o.QueryTable(&Tag{}).Filter("name_in", oldtags).Update(orm.Params{"count": orm.ColValue(orm.ColMinus, 1)})
		o.QueryTable(&TagPost{}).Filter("postid", m.Id).Delete()
	}
	if _, err := orm.NewOrm().Delete(m); err != nil {
		return err
	}
	return nil

}

/** 查询数据 */
func (m *Post) Query() orm.QuerySeter {
	return orm.NewOrm().QueryTable(m)
}
