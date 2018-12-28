package admin

import (
	"myblog/models"
	"time"
)

type MoodController struct {
	baseController
}

func (this *MoodController) Add()  {

	if this.Ctx.Request.Method == "POST"{
		cover := this.GetString("cover")
		content := this.GetString("content")

		var mood models.Mood
		mood.Cover = cover
		mood.Content = content
		mood.Posttime = time.Now()
		if  mood.Insert() == nil{
			this.ShowMsg(200,"发布成功")
		}else {
			this.ShowMsg(201,"发布失败")
		}
	}
	this.display("mood_add")
}

func (this *MoodController) List()  {
	var mood models.Mood
	var list []*models.Mood
	mood.Query().All(&list)
	this.Data["list"] = list
	this.display("mood_list")
}

func (this *MoodController) Delete()  {
	id, _ := this.GetInt64("id")
	var mood models.Mood
	var re  error
	mood.Id = id
	if mood.Read() == nil {
		re = mood.Delete()
	}
	if re == nil {
		this.ShowMsg(200,"删除成功")
	}else {
		this.ShowMsg(201,"删除失败")
	}
}