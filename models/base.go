package models

import (
	"crypto/md5"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"net/url"
	"strings"
)

func init() {
	dbhost := beego.AppConfig.String("dbhost")
	dbport := beego.AppConfig.String("dbport")
	dbuser := beego.AppConfig.String("dbuser")
	dbpassword := beego.AppConfig.String("dbpassword")
	dbname := beego.AppConfig.String("dbname")
	creatTable := beego.AppConfig.DefaultBool("creatTable", false)
	if dbport == "" {
		dbport = "3306"
	}
	dburl := dbuser + ":" + dbpassword + "@tcp(" + dbhost + ":" + dbport + ")/" + dbname + "?charset=utf8"
	orm.RegisterDataBase("default", "mysql", dburl)
	orm.RegisterModel(new(User), new(Option), new(Post), new(Mood), new(Album), new(Tag), new(TagPost))
	orm.RunSyncdb("default", creatTable, true)
	if beego.AppConfig.String("rummode") == "dev" {
		orm.Debug = true
	}
}

func TableName(str string) string {
	return fmt.Sprintf("%s%s", beego.AppConfig.String("dbprefix"), str)
}

func Rawurlencode(str string) string {
	return strings.Replace(url.QueryEscape(str), "+", "%20", -1)
}

func Md5(buf []byte) string {
	hash := md5.New()
	hash.Write(buf)
	return fmt.Sprintf("%x", hash.Sum(nil))
}
