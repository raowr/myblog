package blog

import (
	"myblog/models"
	"strconv"
)

type MainController struct{
	baseController
}

//首页, 只显示前N条
func (this *MainController) Index(){
	var list []*models.Post
	query := new(models.Post).Query().Filter("status",0).Filter("urltype",0)
	count,_ := query.Count()
	if count > 0{
		query.OrderBy("-istop","-views").Limit(this.pagesize,(this.page-1)*this.pagesize).All(&list)
	}
	//this.Data["json"]=list
	//this.ServeJSON()
	this.Data["list"]=list
	this.Data["pagebar"] = models.NewPager(int64(this.page), int64(count), int64(this.pagesize),"/index%d.html").ToString()
	this.setHeadMetas()
	this.display("index")
}

//关于我
func (this *MainController) Introduce()  {
	this.right = ""
	this.display("introduce")
}

//成长记录
func (this *MainController) BlogList()  {
	var list []*models.Post
	query :=new(models.Post).Query().Filter("status",0).Filter("urltype",0)
	count, _ := query.Count()
	if count > 0{
		query.OrderBy("-istop","-posttime").Limit(this.pagesize,(this.page-1)*this.pagesize).All(&list)
	}
	this.Data["list"] = list
	this.Data["pagebar"] = models.NewPager(int64(this.page),int64(count),int64(this.pagesize),"/life%d.html").ToString()
	this.setHeadMetas("成长录")
	this.display("life")
}

func (this *MainController) Book(){
	this.setHeadMetas("留言板")
	this.right = "about.html"
	this.display("book")
}

func (this *MainController) Mood()  {
	var list []*models.Mood
	query := new(models.Mood).Query()
	count, _ := query.Count()
	if count > 0 {
		query.OrderBy("-posttime").Limit(this.pagesize,(this.page-1)*this.pagesize).All(&list)
	}
	this.Data["list"] = list
	this.setHeadMetas("碎言碎语")
	this.right = ""
	this.Data["pagebar"] = models.NewPager(int64(this.page), int64(count),int64(this.pagesize),"/mood%d.html").ToString()
	this.display("mood")
}

func (this *MainController) Album()  {
	pagesize, _ := strconv.Atoi(GetOptionOne("albumsize"))
	if pagesize < 1 {
		pagesize = 12
	}
	var list []*models.Album
	query := new(models.Album).Query().Filter("ishide",0)
	count, _ := query.Count()
	if count > 0 {
		query.OrderBy("-rank","-posttime").Limit(pagesize,(this.page-1)*pagesize).All(&list)
	}
	this.setHeadMetas("光影瞬间")
	this.right = ""
	this.Data["list"] = list
	this.Data["pagebar"] = models.NewPager(int64(this.page),int64(count),int64(pagesize),"/album%d.html").ToString()
	this.display("album")
}