package admin

import (
	"myblog/models"
	"os"
	"runtime"
)

type IndexController struct{
	baseController
}

func (this *IndexController) Index(){
	this.TplName="admin/index.html"
}

func (this *IndexController) Start() {
	this.Data["hostname"],_ = os.Hostname()
	this.Data["gover"] = runtime.Version()
	this.Data["os"] = runtime.GOOS
	this.Data["cpunum"] = runtime.NumCPU()
	this.Data["arch"] = runtime.GOARCH
	this.Data["postnum"],_ = new (models.Post).Query().Count()
	//this.Data["tagnum"] = new(models.Tag).Query().Count()
	this.Data["usernum"],_ = new(models.User).Query().Count()
	this.display("start")
}