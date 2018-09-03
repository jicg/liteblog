package controllers

import (
	"github.com/astaxie/beego"
	"github.com/jicg/liteblog/models"
	"errors"
	"github.com/jicg/liteblog/syserrors"
	"github.com/satori/go.uuid"
)

const SESSION_USER_KEY = "SESSION_USER_KEY"

type NestPreparer interface {
	NestPrepare()
}
type BaseController struct {
	beego.Controller
	IsLogin bool
	User    models.User
	Dao     *models.DB
}

func (ctx *BaseController) Prepare() {
	ctx.Data["Path"] = ctx.Ctx.Request.RequestURI
	ctx.Dao = models.NewDB()
	// 验证用户是否登陆
	ctx.IsLogin = false
	if u, ok := ctx.GetSession(SESSION_USER_KEY).(models.User); ok {
		ctx.User = u
		ctx.Data["User"] = u
		ctx.IsLogin = true
	}
	ctx.Data["IsLogin"] = ctx.IsLogin
	//判断子controller是否实现接口 NestPreparer
	if app, ok := ctx.AppController.(NestPreparer); ok {
		app.NestPrepare()
	}
}

func (ctx *BaseController) MustLogin() {
	if !ctx.IsLogin {
		ctx.Abort500(syserrors.NoUserError{})
	}
}

func (c *BaseController) GetMustString(key string, msg string) string {
	email := c.GetString(key, "")
	if len(email) == 0 {
		c.Abort500(errors.New(msg))
	}
	return email
}

func (ctx *BaseController) Abort500(err error) {
	ctx.Data["error"] = err
	ctx.Abort("500")
}

type H map[string]interface{}

type ResultJsonValue struct {
	Code   int         `json:"code"`
	Msg    string      `json:"msg"`
	Action string      `json:"action,omitempty"`
	Count  int         `json:"count,omitempty"`
	Data   interface{} `json:"data,omitempty"`
}

func (ctx *BaseController) JSONOk(msg string, actions ...string) {
	var action string
	if len(actions) > 0 {
		action = actions[0]
	}
	ctx.Data["json"] = &ResultJsonValue{
		Code:   0,
		Msg:    msg,
		Action: action,
	}
	ctx.ServeJSON()
}

func (ctx *BaseController) JSONOkH(msg string, maps H) {
	if maps == nil {
		maps = H{}
	}
	maps["code"] = 0
	maps["msg"] = msg
	ctx.Data["json"] = maps
	ctx.ServeJSON()
}

func (ctx *BaseController) JSONOkData(count int, data interface{}) {
	ctx.Data["json"] = &ResultJsonValue{
		Code:  0,
		Count: count,
		Msg:   "成功！",
		Data:  data,
	}
	ctx.ServeJSON()
}

func (this *BaseController) UUID() string {
	u,err:=uuid.NewV4()
	if err!=nil{
		this.Abort500(syserrors.NewError("系统错误",err))
	}
	return u.String()
}