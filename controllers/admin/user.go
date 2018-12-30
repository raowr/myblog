package admin

import (
	"fmt"
	"github.com/astaxie/beego/validation"
	"myblog/models"
)

type UserController struct {
	baseController
}

func (this *UserController) Add() {
	if this.Ctx.Request.Method == "POST" {
		usename := this.GetString("username")
		password := this.GetString("password")
		checkPass := this.GetString("checkPass")
		email := this.GetString("email")
		active, _ := this.GetInt64("active")

		valid := validation.Validation{}
		if v := valid.Required(usename, "usename"); !v.Ok {
			this.ShowMsg(201, "请输入用户名")
		} else if v := valid.MaxSize(usename, 15, "usename"); !v.Ok {
			this.ShowMsg(202, "用户名长度不能大于15个字符")
		}

		if v := valid.Required(password, "password"); !v.Ok {
			this.ShowMsg(204, "请输入密码")
		}
		if v := valid.Required(checkPass, "checkPass"); !v.Ok {
			this.ShowMsg(205, "请再次输入密码")
		} else if password != checkPass {
			this.ShowMsg(206, "两次输入的密码不一致")
		}
		if v := valid.Required(email, "email"); !v.Ok {
			this.ShowMsg(207, "请输入email地址")
		} else if v := valid.Email(email, "email"); !v.Ok {
			this.ShowMsg(208, "Email无效")
		}

		if active != 0 {
			active = 1
		}

		user := new(models.User)
		user.Username = usename
		user.Password = models.Md5([]byte(password))
		user.Email = email
		user.Active = int8(active)
		if err := user.Insert(); err != nil {
			this.ShowMsg(209, err.Error())
		}
		this.Redirect("/admin/user/list", 302)
	}
	this.display("user_add")
}

func (this *UserController) Edit() {
	id, _ := this.GetInt64("id")
	user := new(models.User)
	user.Id = id
	fmt.Println(id)
	if user.Read() != nil {
		this.ShowMsg(201, "用户不存在")
	}
	if this.Ctx.Request.Method == "POST" {
		username := this.GetString("username")
		password := this.GetString("password")
		password2 := this.GetString("checkPass")
		email := this.GetString("email")
		active, _ := this.GetInt64("active")

		valid := validation.Validation{}
		if v := valid.Required(username, "username"); !v.Ok {
			this.ShowMsg(201, "请输入用户名")
		} else if v := valid.MaxSize(username, 15, "username"); !v.Ok {
			this.ShowMsg(202, "用户名不能超过15个字符")
		}
		if password != "" {
			if v := valid.Required(password2, "password2"); !v.Ok {
				this.ShowMsg(204, "再次输入密码")
			} else if password != password2 {
				this.ShowMsg(205, "两次输入的密码不一致")
			} else {
				user.Password = models.Md5([]byte(password))
			}
		}
		if v := valid.Required(email, "email"); !v.Ok {
			this.ShowMsg(207, "请输入email地址")
		} else if v := valid.Email(email, "email"); !v.Ok {
			this.ShowMsg(208, "Email无效")
		}

		if active != 0 {
			active = 1
		}
		user.Username = username
		user.Email = email
		user.Active = int8(active)
		user.Update()
		this.ShowMsg(200, "修改成功")
	}
	this.Data["user"] = &user
	this.display("user_edit")
}
func (this *UserController) List() {
	user := new(models.User)
	var list []*models.User
	user.Query().All(&list)
	this.Data["list"] = list
	this.display("user_list")
}

func (this *UserController) Delete() {
	id, _ := this.GetInt64("id")
	user := models.User{Id: id}
	if id == 1 {
		this.ShowMsg(201, "不能删除ID为1的用户")
	}
	if user.Read() == nil {
		user.Delete()
		this.ShowMsg(200, "删除成功")
	} else {
		this.ShowMsg(204, "删除失败")
	}
}
