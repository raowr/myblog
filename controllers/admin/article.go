package admin

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	"myblog/models"
	"strconv"
	"strings"
	"time"
)

type ArticleController struct {
	baseController
}

func (this *ArticleController) Add() {
	this.display("article_add")
}

func (this *ArticleController) Edit() {
	id, _ := this.GetInt64("id")
	post := models.Post{Id: id}
	if post.Read() != nil {
		this.Abort("404")
	}
	post.Tags = strings.Trim(post.Tags, ",")
	this.Data["post"] = post
	this.Data["posttime"] = post.Posttime.Format("2006-01-02 15:04:05")
	this.display("article_edit")
}

func (this *ArticleController) Delete() {
	id, _ := this.GetInt64("id")
	post := models.Post{Id: id}
	if post.Read() != nil {
		post.Delete()
	}
}

func (this *ArticleController) Save() {
	var (
		id      int64  = 0
		title   string = strings.TrimSpace(this.GetString("title"))
		cover   string = strings.TrimSpace(this.GetString("cover"))
		tags    string = strings.TrimSpace(this.GetString("tags"))
		urlname string = strings.TrimSpace(this.GetString("urlname"))
		timestr string = strings.TrimSpace(this.GetString("posttime"))
		content string = this.GetString("content")
		status  int64  = 0
		urltype int8   = 0
		istop   int8   = 0
		post    models.Post
	)
	if title == "" {
		this.ShowMsg(201, "title为空了")
	}
	id, _ = this.GetInt64("id")
	status, _ = this.GetInt64("status")
	if ok, _ := this.GetBool("istop", false); ok {
		istop = 1
	}
	if this.GetString("urltype") == "true" {
		urltype = 1
	}

	if status != 1 && status != 2 {
		status = 0
	}

	if cover == "" {
		cover = "/static/upload/defaultcover.png"
	}

	addtags := make([]string, 0)
	if tags != "" {
		tagarr := strings.Split(tags, ",")
		for _, v := range tagarr {
			if tag := strings.TrimSpace(v); tag != "" {
				exitsts := false
				for _, vv := range addtags {
					if vv == tag {
						exitsts = true
						break
					}
				}
				if !exitsts {
					addtags = append(addtags, tag)
				}
			}
		}
	}
	if id < 1 {
		post.Userid = this.userid
		post.Author = this.username
		post.Posttime = this.getTime()
		post.Updated = this.getTime()
		post.Insert()
		models.Cache.Delete("latestblog")
	} else {
		post.Id = id
		if post.Read() != nil {
			goto RD
		}
		if post.Tags != "" {
			var tagobj models.Tag
			var tagpostobj models.TagPost
			olgtags := strings.Split(strings.Trim(post.Tags, ","), ",")
			//标签统计 -1
			tagobj.Query().Filter("name__in", olgtags).Update(orm.Params{"count": orm.ColValue(orm.ColMinus, 1)})
			tagpostobj.Query().Filter("postid", post.Id).Delete()
		}
	}
	if len(addtags) > 0 {
		for _, v := range addtags {
			tag := models.Tag{Name: v}
			if tag.Read("Name") == orm.ErrNoRows {
				tag.Count = 1
				tag.Insert()
			} else {
				tag.Count += 1
				tag.Update("Count")
			}
			tp := models.TagPost{Tagid: tag.Id, Postid: post.Id, Poststatus: int8(status), Posttime: this.getTime()}
			tp.Insert()
		}
		post.Tags = "," + strings.Join(addtags, ",") + ","
	}
	if posttime, err := time.Parse("2018-01-01 01:01:01", timestr); err == nil {
		post.Posttime = posttime
	} else {
		post.Posttime, _ = time.Parse("2018-01-01 01:01:01", post.Posttime.Format("2018-01-01 01:01:01"))
	}
	post.Status = int8(status)
	post.Title = title
	post.Color = ""
	post.Istop = istop
	post.Cover = cover
	post.Content = content
	post.Brief = string([]rune(content)[:200])
	post.Urlname = urlname
	post.Urltype = urltype
	post.Updated = this.getTime()
	post.Update("tags", "status", "title", "color", "cover", "istop", "content", "urlname", "urltype", "updated", "posttime", "brief")

RD:
	this.Redirect("/admin/article/list", 302)
}

func (this *ArticleController) List() {
	var (
		status     int64
		searchtype string
		keyword    string
		list       []*models.Post
		post       models.Post
	)
	status, _ = this.GetInt64("status", 0)
	query := post.Query().Filter("status", status)
	searchtype = this.GetString("searchtype")
	keyword = this.GetString("keyword")
	if keyword != "" {
		switch searchtype {
		case "title":
			query = query.Filter("title__icontains", keyword)
		case "author":
			query = query.Filter("author__icontains", keyword)
		case "tag":
			query = query.Filter("tag__icontains", keyword)
		}
	}
	count, _ := query.Count()
	if count > 0 {
		query.OrderBy("-istop", "-posttime").All(&list)
	}
	this.Data["searchtype"] = searchtype
	this.Data["keyword"] = keyword
	this.Data["status"] = status
	this.Data["count"] = count
	this.Data["count_0"], _ = post.Query().Filter("status", 0).Count()
	this.Data["count_1"], _ = post.Query().Filter("status", 1).Count()
	this.Data["count_2"], _ = post.Query().Filter("status", 2).Count()
	this.Data["list"] = list
	this.display("article_list")
}

//批处理
func (this *ArticleController) Batch() {
	ids := this.GetStrings("ids")
	op := this.GetString("op")
	var idArr []int64
	for _, v := range ids {
		if id, _ := strconv.Atoi(v); id > 0 {
			idArr = append(idArr, int64(id))
		}
	}

	var post models.Post
	var re int64
	switch op {
	case "topub":
		re, _ = post.Query().Filter("id__in", idArr).Update(orm.Params{"status": 0})
	case "todrafts":
		re, _ = post.Query().Filter("id__in", idArr).Update(orm.Params{"status": 1})
	case "totrash":
		re, _ = post.Query().Filter("id__in", idArr).Update(orm.Params{"status": 2})
	case "delete":
		for _, id := range idArr {
			obj := models.Post{Id: id}
			if obj.Read() != nil {
				obj.Delete()
			}
		}
	}
	if re > 0 {
		this.ShowMsg(200, "更新成功")
	} else {
		this.ShowMsg(201, "更新失敗")
	}
}

//文件上传
func (this *ArticleController) Upload() {
	var coverFile string
	f, h, err := this.GetFile("coverFile")
	if err != nil {
		fmt.Println("getfile err ", err)
	} else {
		this.SaveToFile("coverFile", "static/upload/"+h.Filename)
	}
	if h.Filename != "" {
		coverFile = "static/upload/" + h.Filename
	} else {
		coverFile = "static/upload/defaultcover.png"
	}
	defer f.Close()
	this.print_r(coverFile)
}
