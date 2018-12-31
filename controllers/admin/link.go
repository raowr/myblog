package admin

import (
	"myblog/models"
	"strings"
)

type LinkController struct {
	baseController
}

func (this *LinkController) Add() {
	if this.Ctx.Request.Method == "POST" {
		sitename := this.GetString("sitename")
		url := this.GetString("url")
		rank, _ := this.GetInt8("rank")
		link := new(models.Link)

		link.Sitename = sitename
		link.Url = url
		link.Rank = rank
		if link.Insert() == nil {
			this.ShowMsg(200, "添加成功")
		} else {
			this.ShowMsg(201, "添加失败")
		}
	}
	this.display("link_add")
}

func (this *LinkController) List() {
	var list []*models.Link
	link := new(models.Link)
	if ct, _ := link.Query().Count(); ct > 0 {
		link.Query().All(&list)
	}
	this.Data["list"] = list
	this.display("link_list")
}

func (this *LinkController) Edit() {
	id, _ := this.GetInt64("id")
	link := new(models.Link)
	link.Id = id
	if link.Read() != nil {
		this.ShowMsg(201, "友链不存在")
	}
	if this.Ctx.Request.Method == "POST" {
		siteName := strings.TrimSpace(this.GetString("sitename"))
		url := strings.TrimSpace(this.GetString("url"))
		rank, _ := this.GetInt8("rank")
		link.Sitename = siteName
		link.Url = url
		link.Rank = rank
		if link.Update() == nil {
			this.ShowMsg(200, "修改成功")
		} else {
			this.ShowMsg(201, "修改失败")
		}
	}
	this.Data["link"] = link
	this.display("link_edit")
}

func (this *LinkController) Delete() {
	id, _ := this.GetInt64("id")
	link := new(models.Link)
	link.Id = id
	if link.Read() != nil {
		this.ShowMsg(201, "为找到友链")
	} else {
		link.Delete()
		this.ShowMsg(200, "删除成功")
	}
}
