package controllers

import (
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &Gate{})
}

type Gate struct {
	beego.Controller
}

func (this *Gate) Get() {
	this.TplNames = "gate.html"
}
