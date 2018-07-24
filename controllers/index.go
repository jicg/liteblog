package controllers


type IndexController struct {
	BaseController
}

// @router / [get]
func (c *IndexController) Get() {
	c.TplName = "index.html"
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
