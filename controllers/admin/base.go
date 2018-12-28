package admin

import (
	"github.com/astaxie/beego"
	"myblog/controllers/blog"
	"strconv"
	"strings"
	"time"
)


type baseController struct{
	beego.Controller
	usserid			int64
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
