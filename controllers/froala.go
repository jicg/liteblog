package controllers

import (
	"github.com/astaxie/beego"
	"os"
	"path"
	"strconv"
	"time"
)

type FroalaController struct {
	BaseController
	ImagesFilePath string
	VideosFilePath string
	FilesFilePath  string
}

type UploadimgFroala struct {
	url      string
	title    string
	original string
	state    string
}

const (
	PATH_IMAGE = "assert/images"
	PATH_VIDEO = "assert/videos"
	PATH_FILE  = "assert/files"
)

func (ctx *FroalaController) NestPrepare() {
	if !ctx.IsLogin {
		ctx.Data["json"] = map[string]interface{}{
			"state":    "ERROR",
			"link":     "",
			"title":    "",
			"original": "",
		}
		ctx.ServeJSON()
		ctx.StopRun()
	}
	year, month, _ := time.Now().Date()
	ctx.ImagesFilePath = PATH_IMAGE + "/" + strconv.Itoa(year) + month.String() + "/"
	if err := os.MkdirAll(ctx.ImagesFilePath, 0777); err != nil {
		beego.Error(err)
	}
	ctx.VideosFilePath = PATH_VIDEO + "/" + strconv.Itoa(year) + month.String() + "/"
	if err := os.MkdirAll(ctx.VideosFilePath, 0777); err != nil {
		beego.Error(err)
	}
	ctx.FilesFilePath = PATH_FILE + "/" + strconv.Itoa(year) + month.String() + "/"
	if err := os.MkdirAll(ctx.FilesFilePath, 0777); err != nil {
		beego.Error(err)
	}
}

//添加文章里的图片上传
// @router /uploadimg [post]
func (c *FroalaController) UploadImg() {
	_, h, err := c.GetFile("file")
	if err != nil {
		beego.Error(err)
	}
	fileSuffix := path.Ext(h.Filename)
	newname := strconv.FormatInt(time.Now().UnixNano(), 10) + fileSuffix
	var path string
	if h != nil {
		//保存附件
		path = c.ImagesFilePath + newname
		err = c.SaveToFile("file", path)
		if err != nil {
			beego.Error(err)
		}
		c.Data["json"] = map[string]interface{}{"state": "SUCCESS", "link": "/" + path}
		c.ServeJSON()
	} else {
		c.Data["json"] = map[string]interface{}{"state": "ERROR"}
		c.ServeJSON()
	}
}



//添加文章里的视频上传
// @router /uploadvideo [post]
func (c *FroalaController) UploadVideo() {
	_, h, err := c.GetFile("file")
	if err != nil {
		beego.Error(err)
	}
	fileSuffix := path.Ext(h.Filename)
	newname := strconv.FormatInt(time.Now().UnixNano(), 10) + fileSuffix

	var path string
	if h != nil {
		//保存附件
		path = c.VideosFilePath + newname
		err = c.SaveToFile("file", path) //.Join("attachment", attachment)) //存文件    WaterMark(path)    //给文件加水印
		if err != nil {
			beego.Error(err)
		}
		c.Data["json"] = map[string]interface{}{"state": "SUCCESS", "link": "/" + path}
		c.ServeJSON()
	} else {
		c.Data["json"] = map[string]interface{}{"state": "ERROR"}
		c.ServeJSON()
	}
}

// @router /uploadfile [post]
func (c *FroalaController) UploadFile() {
	_, h, err := c.GetFile("file")
	if err != nil {
		beego.Error(err)
	}
	fileSuffix := path.Ext(h.Filename)
	newname := strconv.FormatInt(time.Now().UnixNano(), 10) + fileSuffix

	var path string
	if h != nil {
		//保存附件
		path = c.FilesFilePath + newname
		err = c.SaveToFile("file", path) //.Join("attachment", attachment)) //存文件    WaterMark(path)    //给文件加水印
		if err != nil {
			beego.Error(err)
		}
		c.Data["json"] = map[string]interface{}{"state": "SUCCESS", "link": "/" + path}
		c.ServeJSON()
	} else {
		c.Data["json"] = map[string]interface{}{"state": "ERROR"}
		c.ServeJSON()
	}
}


////添加wiki里的图片上传
//func (c *FroalaController) UploadWikiImg() {
//	//保存上传的图片
//	_, h, err := c.GetFile("file")
//	if err != nil {
//		beego.Error(err)
//	}
//	var filesize int64
//	fileSuffix := path.Ext(h.Filename)
//	newname := strconv.FormatInt(time.Now().UnixNano(), 10) + fileSuffix // + "_" + filename
//	year, month, _ := time.Now().Date()
//	err = os.MkdirAll(".\\attachment\\wiki\\"+strconv.Itoa(year)+month.String()+"\\", 0777) //..代表本当前exe文件目录的上级，.表示当前目录，没有.表示盘的根目录
//	if err != nil {
//		beego.Error(err)
//	}
//	path1 := ".\\attachment\\wiki\\" + strconv.Itoa(year) + month.String() + "\\" + newname //h.Filename
//	Url := "/attachment/wiki/" + strconv.Itoa(year) + month.String() + "/"
//	err = c.SaveToFile("file", path1) //.Join("attachment", attachment)) //存文件    WaterMark(path)    //给文件加水印
//	if err != nil {
//		beego.Error(err)
//	}
//	filesize, _ = FileSize(path1)
//	filesize = filesize / 1000.0
//	c.Data["json"] = map[string]interface{}{"state": "SUCCESS", "link": Url + newname, "title": "111", "original": "demo.jpg"}
//	c.ServeJSON()
//}