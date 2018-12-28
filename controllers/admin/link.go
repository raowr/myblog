package admin

type LinkController struct {
	baseController
}

func (this *LinkController) Add() {
	this.display("link_add")
}

func (this *LinkController) List()  {
	this.display("link_list")
}