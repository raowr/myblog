package admin

import (
	"myblog/models"
	"time"
)

type AlbumController struct {
	baseController
}

func (this *AlbumController) Add() {
	if this.Ctx.Request.Method == "POST"{
		name := this.GetString("name")
		cover := this.GetString("imageUrl")
		rank,_ := this.GetInt8("rank")

		album := new(models.Album)
		album.Name = name
		album.Cover = cover
		album.Rank = rank
		album.Posttime = time.Now()
		if album.Insert() == nil{
			this.ShowMsg(200,"添加成功")
		}else {
			this.ShowMsg(201,"添加失败")
		}
	}

	this.display("album_add")
}

func (this *AlbumController) Edit()  {
	id,_ := this.GetInt64("id")
	album := models.Album{Id:id}
	if album.Read() != nil {
		this.ShowMsg(201,"相册不存在")
	}
	if this.Ctx.Request.Method == "POST"{
		name := this.GetString("name")
		cover := this.GetString("cover")
		rank,_ := this.GetInt8("rank")
		album.Name = name
		album.Cover = cover
		album.Rank = rank
		if album.Update() == nil{
			this.ShowMsg(200,"修改成功")
		}else {
			this.ShowMsg(201,"修改失败")
		}
	}
	this.Data["data"] = album
	this.display("album_edit")
}

func (this *AlbumController) Delete()  {
	id,_ := this.GetInt64("id")
	album := models.Album{Id:id}
	if album.Read() == nil{
		album.Delete()
		this.ShowMsg(200,"删除成功")
	}else {
		this.ShowMsg(201,"相册不存在")
	}
}

func (this *AlbumController) List()  {
	var list []*models.Album
	album := new(models.Album)
	album.Query().All(&list)
	this.Data["list"] = list
	this.display("album_list")
}


