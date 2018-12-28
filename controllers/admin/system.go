package admin

import (
	"encoding/json"
	"myblog/models"
)

type SystemController struct {
	baseController
}

func (this *SystemController) Setting()  {
	var result []*models.Option
	new(models.Option).Query().All(&result)

	options := make(map[string]string)
	mp := make(map[string]*models.Option)
	for _, v := range result{
		options[v.Name] = v.Value
		mp[v.Name] = v
	}

	if this.Ctx.Request.Method == "POST" {
		var options []*models.Option
		ot := make(map[string]string)
		option := new(models.Option)
		option.Query().All(&options)
		for _, val :=range options{
			ot[val.Name] = val.Value
		}
		type setting struct {
			Sitename 		string `form:"sitename"`
			Subtitle 		string `form:"subtitle"`
			Nickname  		string `form:"nickname"`
			Siteurl 		string `form:"siteurl"`
			Weibo  			string `form:"weibo"`
			Github  		string `form:"github"`
			Pagesize   		string `form:"pagesize"`
			Albumsize   	string `form:"albumsize"`
			Keywords   		string `form:"keywords"`
			Description   	string `form:"description"`
			Theme   		string `form:"theme"`
			Timezone   		string `form:"timezone"`
			Stat  			string `form:"stat"`
		}
		var set setting
		setData := this.Ctx.Input.RequestBody
		json.Unmarshal(setData,&set)
		setMap := map[string]string{
			"Sitename":set.Sitename,
			"Subtitle":set.Subtitle,
			"Nickname":set.Nickname,
			"Siteurl":set.Siteurl,
			"Weibo":set.Weibo,
			"Github":set.Github,
			"Pagesize":set.Pagesize,
			"Albumsize":set.Albumsize,
			"Keywords":set.Keywords,
			"Description":set.Description,
			"Theme":set.Theme,
			"Timezone":set.Timezone,
			"Stat":set.Stat,

		}
		for key, val := range setMap {
			if _, ok := ot[key]; !ok{
				option := new(models.Option)
				option.Name = key
				option.Value = val
				option.Insert()
			}else {
				option := mp[key]
				option.Value = val
				option.Update("Value")
			}
		}
		this.Redirect("/admin/system/setting",302)
	}

	this.Data["now"] = this.getTime()
	this.Data["options"] = options
	this.display("system_setting")
}
