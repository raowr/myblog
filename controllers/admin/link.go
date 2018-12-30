package admin

import "myblog/models"

type LinkController struct {
	baseController
}

func (this *LinkController) Add() {
	if this.Ctx.Request.Method == "POST" {
		sitename := this.GetString("sitename")
		url := this.GetString("url")
		rank,_ := this.GetInt8("rank")
		link := new(models.Link)

		link.Sitename = sitename
		link.Url = url
		link.Rank = rank
		if link.Insert() == nil{
			this.ShowMsg(200,"添加成功")
		}else {
			this.ShowMsg(201,"添加失败")
		}
	}
	this.display("link_add")
}

func (this *LinkController) List()  {
	var list []*models.Link
	link := new(models.Link)
	if ct, _ := link.Query().Count();ct > 0{
		link.Query().All(&list)
	}
	this.Data["list"] = list
	this.display("link_list")
}