package routers

import (
	"github.com/jicg/liteblog/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.ErrorController(&controllers.ErrorController{})
	beego.Include(&controllers.IndexController{})
	//beego.Router("/", &controllers.MainController{})
	//beego.Router("/about", &controllers.MainController{})
	//beego.Router("/message", &controllers.MainController{})
}
