package controllers

import (
)

type IndexController struct {
	BaseController
}

// @router / [get]
func (c *IndexController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.html"
}


// @router /message [get]
func (c *IndexController) GetMessage() {
	c.TplName = "message.html"
}


// @router /about [get]
func (c *IndexController) GetAbout() {
	c.TplName = "about.html"
}
