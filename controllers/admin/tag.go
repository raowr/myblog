package admin

import (
	"myblog/models"
	"strconv"
	"strings"
)

type TagController struct {
	baseController
}

func (this *TagController) Index() {

	var tag models.Tag
	var list []*models.Tag
	tag.Query().All(&list)
	this.Data["list"] = list
	this.display("tag_list")
}

func (this *TagController) Bacth()  {
	ids := this.GetStrings("ids")
	op := this.GetString("op")

	idarr := make([]int64,0)
	for _, v := range ids{
		if id, _ := strconv.Atoi(v); id > 0{
			idarr = append(idarr,int64(id))
		}
	}

	switch op {
	case "upcount":
		for _, id:= range idarr{
			obj := models.Tag{Id:id}
			if obj.Read() != nil{
				obj.UpCount()
			}
		}
	case "merge":
		toname := strings.TrimSpace(this.GetString("toname"))
		if toname != "" && len(idarr) > 0{
			tag := new(models.Tag)
			tag.Name = toname
			if tag.Read("name") != nil{
				tag.Count = 0
				tag.Insert()
			}
			for _, id := range idarr{
				obj := models.Tag{Id:id}
				if obj.Read() == nil{
					obj.MergeTo(tag)
					obj.Delete()
				}
			}
			tag.UpCount()
		}
	case "delete":
		for _, id := range idarr{
			obj := models.Tag{Id:id}
			if obj.Read() == nil{
				obj.Delete()
			}
		}

	}
	this.ShowMsg(200,"操作成功")
}
