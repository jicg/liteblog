package main

import (
	_ "github.com/jicg/liteblog/routers"
	"github.com/astaxie/beego"
)

func main() {
	initSession()
	beego.Run()
}

func initSession() {
	//gob.Register(&models.User{})
	//https://beego.me/docs/mvc/controller/session.md
	beego.BConfig.WebConfig.Session.SessionOn = true
	beego.BConfig.WebConfig.Session.SessionName = "liteblog-key"
	beego.BConfig.WebConfig.Session.SessionProvider = "file"
	beego.BConfig.WebConfig.Session.SessionProviderConfig = "data/session"
}
