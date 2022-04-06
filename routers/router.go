package routers

import (
	"github.com/beego/beego/v2/server/web"
	"github.com/jicg/liteblog/controllers"
)

func init() {
	web.ErrorController(&controllers.ErrorController{})
	web.Include(
		&controllers.IndexController{},
		&controllers.UserController{},
	)
	web.AddNamespace(
		web.NewNamespace(
			"note",
			web.NSInclude(&controllers.NoteController{}),
		),
		web.NewNamespace(
			"upload",
			web.NSInclude(&controllers.UploadController{}),
		),
		web.NewNamespace(
			"praise",
			web.NSInclude(&controllers.PraiseController{}),
		),
		web.NewNamespace(
			"message",
			web.NSInclude(&controllers.MessageController{}),
		),
	)
}
