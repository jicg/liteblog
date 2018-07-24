package controllers

import (
	"fmt"
	"github.com/jicg/liteblog/syserrors"
	"github.com/astaxie/beego/logs"
)

type ErrorController struct {
	BaseController
}

func (c *ErrorController) Error404() {
	c.TplName = "error/404.html"
	if (c.IsAjax()) {
		c.jsonerror(syserrors.Error404{})
	} else {
		c.Data["content"] = "非法访问"
	}
}

func (c *ErrorController) Error500() {
	c.TplName = "error/500.html"
	// 获取c.Data["error"] 错误，默认为 UnKnowError
	var derr error;
	err, ok := c.Data["error"].(error)
	if ok {
		derr = err
	} else {
		derr = syserrors.UnKnowError{}
	}
	// 将error转成 syserrors.Error ，转不了就默认 UnKnowError，
	var dserr syserrors.Error
	if serr, ok := derr.(syserrors.Error); ok {
		dserr = serr
	} else {
		dserr = syserrors.NewError(err.Error(), err)
	}
	//打印日志
	if err := dserr.ReasonError(); err != nil {
		logs.Error(dserr.Error(), err)
	}
	//输出
	if (c.IsAjax()) {
		c.jsonerror(dserr)
	} else {
		c.Data["content"] = fmt.Sprintf("错误：%s", dserr.Error())
	}
}
func (ctx *ErrorController) jsonerror(err syserrors.Error) {
	ctx.Ctx.Output.Status = 200
	ctx.Data["json"] = &ResultJsonValue{
		Code: err.Code(),
		Msg:  err.Error(),
	}
	ctx.ServeJSON()
}
