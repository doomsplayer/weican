package main

import (
	//"fmt"
	"github.com/astaxie/beego"
	"weican/controllers"
)

func main() {
	beego.Router("/", &controllers.MainController{})
	beego.Run()
}
