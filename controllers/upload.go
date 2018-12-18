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
	"encoding/json"
	"io"
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
func (c *UploadController) UploadImg() {

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

//@router /wangeditorfiles [post]
func (c *UploadController) WangeditorUploadFile() {
	hs, err := c.GetFiles("files")
	if err != nil {
		c.JSONOkH("", H{
			"errno": -1,
			"msg":   "上传失败：" + err.Error(),
		})
		return
	}
	var len= len(hs)
	paths := make([]string, len)
	for index, h := range hs {
		fileSuffix := path.Ext(h.Filename)
		newname := strconv.FormatInt(time.Now().UnixNano(), 10) + fileSuffix
		//保存附件
		path := c.FilesFilePath + newname
		err = saveFile(h, path) //.Join("attachment", attachment)) //存文件    WaterMark(path)    //给文件加水印
		if err != nil {
			c.JSONOkH("", H{
				"errno": -1,
				"msg":   "上传失败：" + err.Error(),
			})
		}
		go saveFileinfo(path, h)
		paths[index] = "/" + path
	}
	c.JSONOkH("", H{
		"errno": 0,
		"data":  paths,
	})
}

func saveFile(f *multipart.FileHeader,path string) error {
	file, err := f.Open()
	defer file.Close()
	if err != nil {
		return err
	}
	//create destination file making sure the path is writeable.
	dst, err := os.Create(path)
	defer dst.Close()
	if err != nil {
		return err
	}
	//copy the uploaded file to the destination file
	if _, err := io.Copy(dst, file); err != nil {
		return err
	}
	return nil
}
