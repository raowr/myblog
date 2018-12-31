package admin

import (
	"github.com/astaxie/beego"
	"myblog/controllers/blog"
	"myblog/models"
	"strconv"
	"strings"
	"time"
)


type baseController struct{
	beego.Controller
	userid			int64
	username		string
	moduleName 		string
	controllerName 	string
	actionName 		string
}

func (this *baseController) Prepare(){
	this.moduleName = "admin"
	controllerName, actionName := this.GetControllerAndAction()
	this.controllerName = strings.ToLower(controllerName[0:len(controllerName)-10])
	this.actionName = strings.ToLower(actionName)
	this.auth()
}

/**
登录认证
 */
func (this *baseController) auth()  {
	if this.controllerName == "account" && (this.actionName == "login" || this.actionName == "logout"){

	}else {
		arr := strings.Split(this.Ctx.GetCookie("auth"), "|")
		if len(arr) == 2 {
			idstr,password := arr[0],arr[1]
			userid,_ := strconv.ParseInt(idstr,10,0)
			if userid > 0 {
				var user models.User
				user.Id = userid
				if user.Read() == nil && password == models.Md5([]byte(this.getClientIp() + "|"+user.Password)) {
					this.userid = user.Id
					this.username = user.Username
				}
			}
		}
		if this.userid == 0 {
			this.Redirect("/admin/login",302)
		}

	}
}

func  (this *baseController) display(tpl ...string) {
	var tplName string
	if len(tpl) == 1 {
		tplName = this.moduleName + "/" + tpl[0] +".html"
	}else{
		tplName = this.moduleName + "/" + this.controllerName + "_" + this.actionName + ".html"
	}
	this.Layout = this.moduleName + "/layout.html"
	this.TplName = tplName
}

func (this *baseController) getTime() time.Time {
	option := blog.GetOptionOne("timezone")
	timezone := float64(0)
	if option != ""{
		timezone, _ = strconv.ParseFloat(option, 64)
	}
	add := timezone * float64(time.Hour)
	return time.Now().UTC().Add(time.Duration(add))
}

//权限验证
func (this *baseController) checkPermission() {
	if this.userid != 1 && this.controllerName == "user" {
		this.ShowMsg(201,"抱歉，只有超级管理员才能进行该操作！")
	}
}

func (this *baseController) print_r(options interface{})  {
	this.Data["json"] = options
	this.ServeJSON()
	this.StopRun()
}

func (this *baseController) ShowMsg(code int,msg string) {
	type State struct {
		Code int
		Msg string
	}
	data := State{
		200,
		"sussce",
	}
	if code != 0 {
		data.Code = code
	}
	if msg != "" {
		data.Msg = msg
	}
	this.print_r(data)
}

//获取用户IP地址
func (this *baseController) getClientIp() string {
	s := strings.Split(this.Ctx.Request.RemoteAddr, ":")
	return s[0]
}

func (this *baseController) GetTime() time.Time {
	timezone := float64(0)
	add := timezone * float64(time.Hour)
	return time.Now().UTC().Add(time.Duration(add))
}
