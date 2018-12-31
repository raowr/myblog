package admin

import (
	"myblog/models"
	"strconv"
	"strings"
)

type AccountController struct {
	baseController
}

func (this *AccountController) Login()  {
	if this.Ctx.Request.Method == "POST" {
		account := strings.TrimSpace(this.GetString("account"))
		password := strings.TrimSpace(this.GetString("password"))
		if account != "" && password != ""{
			var user models.User
			user.Username = account
			if user.Read("Username") != nil || user.Password != models.Md5([]byte(password)){
				this.ShowMsg(201,"账号或密码错误")
			}else if user.Active == 0{
				this.ShowMsg(202,"该账号未激活")
			}else {
				user.Logincount += 1
				user.Lastip = this.getClientIp()
				user.Lastlogin = this.getTime()
				user.Update()
				authkey := models.Md5([]byte(this.getClientIp()+"|"+user.Password))
				this.Ctx.SetCookie("auth",strconv.FormatInt(user.Id,10)+"|"+authkey)
			}
			this.Redirect("/admin",302)
		}
	}
	this.TplName = "admin/account_login.html"
}

func (this *AccountController) Logout()  {
	this.Ctx.SetCookie("auth","")
	this.Redirect("/admin/login",302)
}
