package blog

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"myblog/models"
	"strconv"
)

type baseController struct{
	beego.Controller
	options  map[string]string
	right    string
	page     int
	pagesize int
}

func (this *baseController) Prepare() {
	this.right = "right.html"
	this.Data["hidejs"] = `<!--[if lt IE 9]>
  <script src="/views/double/js/modernizr.js"></script>
  <![endif]-->`
	var (
		pagesize int
		err error
		page int
	)
	if page, err = strconv.Atoi(this.Ctx.Input.Param(":page")); err != nil || page < 1 {
		page = 1
	}
	if pagesize, err = strconv.Atoi(GetOptionOne("pagesize")); err != nil || pagesize < 1{
		pagesize = 10
	}
	this.page = page
	this.pagesize = pagesize
}

func (this *baseController) display(tpl string) {
	theme := "blog"
	this.Layout = theme + "/" + "layout.html"
	this.TplName = theme + "/" + tpl+ ".html"
	this.Data["root"] = "/" + beego.AppConfig.String("ViewsPath") + "/" + theme + "/"
	this.LayoutSections = make(map[string]string)
	this.LayoutSections["head"] = theme + "/head.html"
	if tpl == "index" {
		this.LayoutSections["banner"] = theme + "/banner.html"
		this.LayoutSections["middle"] = theme + "/middle.html"
	}
	if this.right != ""{
		this.LayoutSections["right"] = theme + "/" + this.right
	}
	this.LayoutSections["foot"] = theme + "/foot.html"
}

func (this *baseController) setHeadMetas(params ...string) {
	options := getOption()
	if len(params)> 0 {
		this.Data["title"] = params[0]
	}else {
		this.Data["title"] = options["title"]
	}
	this.Data["keywords"] = options["keywords"]
	this.Data["description"] = options["description"]
	this.Data["nickname"] = options["nickname"]
	//this.ServeJSON()

}

func  GetOptionOne(name string) string {
	options := getOption()
	return options[name]
}

func  getOption() map[string]string {
	var result []*models.Option
	o := orm.NewOrm()
	o.QueryTable(&models.Option{}).All(&result)
	options := make(map[string]string)
	for _, v := range result{
		options[v.Name] = v.Value
	}
	return options
}

func (this *baseController) print_r(options interface{})  {
	this.Data["json"] = options
	this.ServeJSON()
	this.StopRun()
}