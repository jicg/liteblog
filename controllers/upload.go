package controllers

import (
	"github.com/astaxie/beego"
	"os"
	"path"
	"strconv"
	"time"
	"github.com/jicg/liteblog/syserrors"
	"mime/multipart"
	"io/ioutil"
	"github.com/gin-gonic/gin/json"
)

type UploadController struct {
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

func (ctx *UploadController) NestPrepare() {
	ctx.MustLogin()
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

//图片上传
// @router /uploadimg [post]
func (c *UploadController) UploadImg(){

	_, h, err := c.GetFile("file")
	if err != nil {
		c.Abort500(err)
	}
	fileSuffix := path.Ext(h.Filename)
	newname := strconv.FormatInt(time.Now().UnixNano(), 10) + fileSuffix
	var path string
	if h != nil {
		//保存附件
		path = c.ImagesFilePath + newname
		err = c.SaveToFile("file", path)
		if err != nil {
			c.Abort500(err)
		}
		go saveFileinfo(path, h)
		c.JSONOkH("上传成功", H{
			"link": "/" + path,
		})
	} else {
		c.Abort500(syserrors.NewError("上传失败", nil))
	}
}

func saveFileinfo(path string, f *multipart.FileHeader) error {
	bs, err := json.Marshal(f)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(path+".json", bs, 0777)
}

// @router /uploadfile [post]
func (c *UploadController) UploadFile() {
	_, h, err := c.GetFile("file")
	if err != nil {
		c.Abort500(err)
	}
	fileSuffix := path.Ext(h.Filename)
	newname := strconv.FormatInt(time.Now().UnixNano(), 10) + fileSuffix
	var path string
	if h != nil {
		//保存附件
		path = c.FilesFilePath + newname
		err = c.SaveToFile("file", path) //.Join("attachment", attachment)) //存文件    WaterMark(path)    //给文件加水印
		if err != nil {
			c.Abort500(err)
		}
		go saveFileinfo(path, h)
		c.JSONOkH("上传成功", H{
			"link": "/" + path,
		})
	} else {
		c.Abort500(syserrors.NewError("上传失败", nil))
	}
}
