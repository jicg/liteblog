package controllers

import (
	"github.com/astaxie/beego"
)

const (
	USER_KEY         = "USER"
	KEY_VERIFICATION = "key_verification"
)


type NestPreparer interface {
	NestPrepare()
}
type BaseController struct {
	beego.Controller
	//User    *models.User
	//IsLogin bool
}


func (ctx *BaseController) Prepare() {
	//user := ctx.GetSession(USER_KEY)
	//ctx.Data["IsLogin"] = false
	//ctx.IsLogin = false
	//if user != nil {
	//	u := user.(*models.User)
	//	if u.Id != 0 {
	//		ctx.User = u
	//		ctx.Data["User"] = ctx.User
	//		ctx.IsLogin = true
	//		ctx.Data["IsLogin"] = true
	//	}
	//}
	ctx.Data["Title"] = "论坛"
	ctx.Data["Page"] = "index"
	if app, ok := ctx.AppController.(NestPreparer); ok {
		app.NestPrepare()
	}
}

