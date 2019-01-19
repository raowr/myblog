package routers

import (
	"github.com/astaxie/beego"
	"myblog/controllers"
	"myblog/controllers/admin"
	"myblog/controllers/blog"
)

func init() {
	// beego.Router("/", &controllers.MainController{})
	beego.Router("/", &blog.MainController{}, "*:Index")
	beego.Router("/index", &blog.MainController{}, "*:Index")
	beego.Router("/user", &controllers.UserController{}, "*:Index")
	beego.Router("/admin", &admin.IndexController{}, "*:Index")
	beego.Router("/admin/login", &admin.AccountController{}, "*:Login")
	beego.Router("/admin/logout", &admin.AccountController{}, "*:Logout")
	beego.Router("/admin/profile", &admin.AccountController{}, "*:Profile")
	beego.Router("/admin/start", &admin.IndexController{}, "*:Start")
	beego.Router("/life:page:int.html", &blog.MainController{}, "*:BlogList")
	beego.Router("/introduce.html", &blog.MainController{}, "*:Introduce")
	beego.Router("/life.html", &blog.MainController{}, "*:BlogList")
	beego.Router("book.html", &blog.MainController{}, "*:Book")
	beego.Router("/mood.html", &blog.MainController{}, "*:Mood")
	beego.Router("mood:page:int.html", &blog.MainController{}, "*:Mood")
	beego.Router("/album.html", &blog.MainController{}, "*:Album")
	beego.Router("/album:page:int.html", &blog.MainController{}, "*:Album")

	//系统管理
	beego.Router("/admin/system/setting", &admin.SystemController{}, "*:Setting")
	beego.Router("/admin/article/add", &admin.ArticleController{}, "*:Add")
	beego.Router("/admin/article/edit", &admin.ArticleController{}, "*:Edit")
	beego.Router("/admin/article/list", &admin.ArticleController{}, "*:List")
	beego.Router("/admin/article/save", &admin.ArticleController{}, "*:Save")
	beego.Router("/admin/article/batch", &admin.ArticleController{}, "*:Batch")

	//说说管理
	beego.Router("/admin/mood/add", &admin.MoodController{}, "*:Add")
	beego.Router("/admin/mood/list", &admin.MoodController{}, "*:List")
	beego.Router("/admin/mood/delete", &admin.MoodController{}, "*:Delete")

	beego.Router("/admin/tag", &admin.TagController{}, "*:Index")
	beego.Router("/admin/tag/batch", &admin.TagController{}, "*:Bacth")

	beego.Router("/admin/user/add", &admin.UserController{}, "*:Add")
	beego.Router("/admin/user/list", &admin.UserController{}, "*:List")
	beego.Router("/admin/user/edit", &admin.UserController{}, "*:Edit")
	beego.Router("/admin/user/delete", &admin.UserController{}, "*:Delete")

	beego.Router("/admin/album/add", &admin.AlbumController{}, "*:Add")
	beego.Router("/admin/album/list", &admin.AlbumController{}, "*:List")
	beego.Router("/admin/album/edit", &admin.AlbumController{}, "*:Edit")
	beego.Router("/admin/album/delete", &admin.AlbumController{}, "*:Delete")

	beego.Router("/admin/link/add", &admin.LinkController{}, "*:Add")
	beego.Router("/admin/link/list", &admin.LinkController{}, "*:List")
	beego.Router("/admin/link/edit", &admin.LinkController{}, "*:Edit")
	beego.Router("/admin/link/delete", &admin.LinkController{}, "*:Delete")

}
