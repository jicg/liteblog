package main

import (
	"encoding/gob"
	"encoding/json"
	"fmt"
	"github.com/beego/beego/v2/core/logs"
	beego "github.com/beego/beego/v2/server/web"
	"github.com/jicg/liteblog/models"
	_ "github.com/jicg/liteblog/models"
	_ "github.com/jicg/liteblog/routers"
	"os"
	"strings"
)

//go:generate bee generate routers -ctrlDir=controllers -routersFile=mygen.go -routersPkg="controllers"

func main() {
	beego.BConfig.WebConfig.CommentRouterPath = "./controllers"
	initLog()
	initSession()
	initTemplate()
	beego.Run()
}
func initLog() {
	if err := os.MkdirAll("data/logs", 0777); err != nil {
		logs.Error(err)
		return
	}
	logs.SetLogger("file", `{"filename":"data/logs/lyblog.log","level":7,"maxlines":1000,"maxsize":100,"daily":true,"maxdays":10}`)
	logs.Async(1e3)
}

func initSession() {
	gob.Register(models.User{})
	//https://beego.vip/docs/mvc/controller/session.md
	beego.SetStaticPath("assert", "assert")
	beego.BConfig.WebConfig.Session.SessionOn = true
	beego.BConfig.WebConfig.Session.SessionName = "liteblog-key"
	beego.BConfig.WebConfig.Session.SessionProvider = "file"
	beego.BConfig.WebConfig.Session.SessionProviderConfig = "data/session"
}
func initTemplate() {
	beego.AddFuncMap("equrl", func(x, y string) bool {
		s1 := strings.Trim(x, "/")
		s2 := strings.Trim(y, "/")
		return strings.Compare(s1, s2) == 0
	})
	beego.AddFuncMap("eq2", func(x, y interface{}) bool {
		s1 := fmt.Sprintf("%v", x)
		s2 := fmt.Sprintf("%v", y)
		return strings.Compare(s1, s2) == 0
	})
	beego.AddFuncMap("add", func(x, y int) int {
		return x + y
	})
	beego.AddFuncMap("json", func(obj interface{}) string {
		bs, err := json.Marshal(obj)
		if err != nil {
			return "{id:0}"
		}
		return string(bs)
	})

}
