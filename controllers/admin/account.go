package admin

import (
	"myblog/models"
	"strconv"
	"strings"
)

type AccountController struct {
	baseController
}

func (this *AccountController) Login() {
	if this.Ctx.Request.Method == "POST" {
		account := strings.TrimSpace(this.GetString("account"))
		password := strings.TrimSpace(this.GetString("password"))
		if account != "" && password != "" {
			var user models.User
			user.Username = account
			if user.Read("Username") != nil || user.Password != models.Md5([]byte(password)) {
				//this.ShowMsg(201,"账号或密码错误")
				re := map[string]interface{}{
					"code": 201,
					"msg":  "账号或密码错误",
				}
				this.Data["json"] = re
				this.ServeJSON()
			} else if user.Active == 0 {
				re := map[string]interface{}{
					"code": 202,
					"msg":  "该账号未激活",
				}
				this.Data["json"] = re
				this.ServeJSON()
				//this.ShowMsg(202,"该账号未激活")
			} else {
				user.Logincount += 1
				user.Lastip = this.getClientIp()
				user.Lastlogin = this.getTime()
				user.Update()
				authkey := models.Md5([]byte(this.getClientIp() + "|" + user.Password))
				this.Ctx.SetCookie("auth", strconv.FormatInt(user.Id, 10)+"|"+authkey)
			}
			this.Redirect("/admin", 302)
		}
	}
	this.TplName = "admin/account_login.html"
}

func (this *AccountController) Logout() {
	this.Ctx.SetCookie("auth", "")
	this.Redirect("/admin/login", 302)
}

func (this *AccountController) Profile() {
	user := models.User{Id: this.userid}
	if err := user.Read(); err != nil {
		this.ShowMsg(201, "未找到信息")
	}
	if this.Ctx.Request.Method == "POST" {
		password := strings.TrimSpace(this.GetString("password"))
		newpassword := strings.TrimSpace(this.GetString("newpassword"))
		newpassword2 := strings.TrimSpace(this.GetString("newpassword2"))
		if newpassword != "" {
			if password == "" || models.Md5([]byte(password)) != user.Password {
				this.ShowMsg(202, "当前密码错误")
			} else if len(newpassword) < 6 {
				this.ShowMsg(204, "密码长度不能少于6个字符")
			} else if newpassword != newpassword2 {
				this.ShowMsg(205, "两次密码不一致")
			} else {
				user.Password = models.Md5([]byte(newpassword))
				if user.Update("password") == nil {
					this.ShowMsg(200, "修改成功")
				} else {
					this.ShowMsg(206, "修改失败")
				}
			}

		}
	}
	this.Data["user"] = user
	this.display("account_profile")
}
