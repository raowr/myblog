package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"myblog/models"
	"time"
)

type UserController struct {
	beego.Controller
}

func (This *UserController) Index() {
	var list []*models.User
	query := new(models.User).Query()
	count, _ := query.Count()
	if count > 0 {
		query.All(&list)
	}else {
		user := &models.User{
			Username:"jams",
			Password:"123456",
			Email   :"jams@qq.com",
			Lastlogin:  time.Now(),
			Logincount: 1,
			Lastip     :"192.168.1.1",
			Authkey   : "123456",
			Active    :int8(2),
		}
		err := user.Insert()
		if err == nil {
			fmt.Println("插入成功")
		}
	}
	fmt.Println("count:",count)
	// users := &models.User{
	// 	Username:"jams1",
	// 	Password:"123456",
	// 	Email   :"jams@qq.com",
	// 	Lastlogin:  time.Now(),
	// 	Logincount: 1,
	// 	Lastip     :"192.168.1.1",
	// 	Authkey   : "123456",
	// 	Active    :int8(2),
	// }
	// This.Data["list"] = &list
	This.Data["json"] = &list
	// This.Data["json"] = &users
	This.ServeJSON()
	This.TplName = "usercontroller/index.html"
}

func (This *UserController) Get() {
	This.Data["json"] = "11"
	This.ServeJSON()
}
