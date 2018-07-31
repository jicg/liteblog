package controllers

import (
	"github.com/go-xweb/uuid"
	"github.com/jicg/liteblog/models"
	"time"
	"github.com/jicg/liteblog/syserrors"
	"github.com/jinzhu/gorm"
	"github.com/PuerkitoBio/goquery"
	"bytes"
)

type NoteController struct {
	BaseController
}

func (ctx *NoteController) NestPrepare() {
	ctx.MustLogin()
	if ctx.User.Role != 0 {
		ctx.Abort500(syserrors.NewError("您没有权限修改文章", nil))
	}
}

// @router /new [get]
func (ctx *NoteController) NewPage() {
	ctx.Data["key"] = uuid.NewUUID().String()
	ctx.TplName = "note_new.html"
}

// @router /edit/:key [get]
func (ctx *NoteController) EditPage() {
	key := ctx.Ctx.Input.Param(":key")
	note, err := ctx.Dao.QueryNoteByKey(key)
	if err != nil {
		ctx.Abort500(syserrors.NewError("文章不存在", err))
	}
	ctx.Data["note"] = note
	ctx.Data["key"] = key
	ctx.TplName = "note_new.html"
}

// @router /del/:key [post]
func (ctx *NoteController) Del() {
	key := ctx.Ctx.Input.Param(":key")
	if err := ctx.Dao.DelNoteByKey(key, int(ctx.User.ID)); err != nil {
		ctx.Abort500(syserrors.NewError("删除失败", err))
	}
	ctx.JSONOk("删除成功！", "/")
}

// @router /save/:key [post]
func (ctx *NoteController) Save() {
	key := ctx.Ctx.Input.Param(":key")
	title := ctx.GetMustString("title", "标题不能为空！")
	content := ctx.GetMustString("content", "内容不能为空！")

	summary, _ := getSummary(content)
	note, err := ctx.Dao.QueryNoteByKeyAndUserId(key, int(ctx.User.ID))
	var n models.Note
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			ctx.Abort500(syserrors.NewError("保存失败！", err))
		}
		n = models.Note{
			Key:     key,
			Summary: summary,
			Title:   title,
			Content: content,
			UserID:  int(ctx.User.ID),
		}
	} else {
		n = note
		n.Title = title
		n.Content = content
		n.Summary = summary
		n.UpdatedAt = time.Now()
	}
	if err := ctx.Dao.SaveNote(&n); err != nil {
		ctx.Abort500(syserrors.NewError("保存失败！", err))
	}
	ctx.JSONOk("成功")
}



func getSummary(content string) (string, error) {
	var buf bytes.Buffer
	buf.Write([]byte(content))
	doc, err := goquery.NewDocumentFromReader(&buf)
	if err != nil {
		return "", err
	}
	str := doc.Find("body").Text()
	if len(str) > 600 {
		str = str[0:600] + "..."
	}
	return str, nil
}
