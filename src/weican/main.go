package main

import (
	"weican/controllers"
	"github.com/astaxie/beego"
	"fmt"
)

func main() {
	beego.Router("/", &controllers.MainController{})
	beego.Run()
	beego.
}

