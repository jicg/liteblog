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

// @router /new [get]
func (ctx *NoteController) NewPage() {
	ctx.MustLogin()
	ctx.Data["key"] = uuid.NewUUID().String()
	ctx.TplName = "note_new.html"
}

// @router /save/:key [post]
func (ctx *NoteController) Save() {
	ctx.MustLogin()
	key := ctx.Ctx.Input.Param(":key")
	title := ctx.GetMustString("title", "标题不能为空！")
	content := ctx.GetMustString("content", "内容不能为空！")

	summary, _ := getSummary(content)
	note, err := models.QueryNoteByKeyAndUserId(key, int(ctx.User.ID))
	var n *models.Note
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			ctx.Abort500(syserrors.NewError("保存失败！", err))
		}
		n = &models.Note{
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
	if err := models.SaveNote(n); err != nil {
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
