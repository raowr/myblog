package admin

type AlbumController struct {
	baseController
}

func (this *AlbumController) Add() {
	this.display("album_add")
}

func (this *AlbumController) List()  {
	this.display("album_list")
}
