package main

import (
	"github.com/astaxie/beego"
	_ "planB/routers"
)

func main() {
	beego.BConfig.WebConfig.Session.SessionOn = true
	beego.Run()
}

