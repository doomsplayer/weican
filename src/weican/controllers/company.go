package controllers

import (
	"github.com/astaxie/beego"
)

func init() {
	beego.Router(`/company/about`, monoCompany, `get:About`)
}

type company struct {
	beego.Controller
}

var monoCompany = &company{}

func (this *company) About() {
	this.TplNames = `company/about.html`
}
