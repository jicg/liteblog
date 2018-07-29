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
	limit := 10;
	page, err := c.GetInt("page", 1)
	if err != nil || page < 1 {
		page = 1;
	}
	title := c.GetString("title", "")
	if ns, _ := models.QueryNotesByPage(page, limit, title); ns != nil {
		c.Data["notes"] = ns
	}
	var totpage int = 0;
	totcnt, _ := models.QueryNotesCount(title)
	if totcnt%limit == 0 {
		totpage = totcnt / limit
	} else {
		totpage = totcnt/limit + 1
	}
	c.Data["totpage"] = totpage
	c.Data["page"] = page
	c.Data["title"] = title
	c.TplName = "index.html"
}

// @router /details/:key [get]
func (c *IndexController) GetDetail() {
	key := c.Ctx.Input.Param(":key")
	note, err := models.QueryNoteByKey(key)
	if err != nil {
		c.Abort500(syserrors.NewError("文章不存在", err))
	}
	go models.AllVisitCount(key)
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
