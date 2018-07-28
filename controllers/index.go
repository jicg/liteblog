package controllers

import (
	"github.com/jicg/liteblog/models"
	"github.com/jicg/liteblog/syserrors"
)

type IndexController struct {
	BaseController
}

// @router / [get]
func (c *IndexController) Get() {
	if ns, _ := models.QueryNotesBy(0, 10); ns != nil {
		c.Data["notes"] = ns
	}
	c.TplName = "index.html"
}

// @router /details/:key [get]
func (c *IndexController) GetDetail() {
	key := c.Ctx.Input.Param(":key")
	note, err := models.QueryNoteByKey(key)
	if err != nil {
		c.Abort500(syserrors.NewError("文章不存在", err))
	}
	c.Data["note"] = note
	c.TplName = "details.html"
}

// @router /comment/:key [get]
func (c *IndexController) GetComment() {
	key := c.Ctx.Input.Param(":key")
	note, err := models.QueryNoteByKey(key)
	if err != nil {
		c.Abort500(syserrors.NewError("文章不存在", err))
	}
	c.Data["note"] = note
	c.TplName = "comment.html"
}

// @router /user [get]
func (c *IndexController) GetUser() {
	c.TplName = "user.html"
}

// @router /reg [get]
func (c *IndexController) GetReg() {
	c.TplName = "reg.html"
}

// @router /message [get]
func (c *IndexController) GetMessage() {
	c.TplName = "message.html"
}

// @router /about [get]
func (c *IndexController) GetAbout() {
	c.TplName = "about.html"
}
