package controllers

import (
	"github.com/jicg/liteblog/models"
	"github.com/jicg/liteblog/syserrors"
)

type MessageController struct {
	BaseController
}

func (ctx *MessageController) NestPrepare() {

}

// @router /new/?:key [post]
func (ctx *MessageController) NewMessage() {
	ctx.MustLogin()
	key :=  ctx.UUID()
	content := ctx.GetMustString("content", "内容不能为空")
	notekey := ctx.Ctx.Input.Param(":key")
	m := &models.Message{
		UserID:  int(ctx.User.ID),
		User:    ctx.User,
		Key:     key,
		NoteKey: notekey,
		Content: content,
	}
	if err := ctx.Dao.SaveMessage(m); err != nil {
		ctx.Abort500(syserrors.NewError("保存失败！", err))
	}
	ctx.JSONOkH("保存成功！", H{
		"data": m,
	})
}

// @router /count [get]
func (ctx *MessageController) Count() {
	count, err := ctx.Dao.QueryMessageForNoteCount("")
	if err != nil {
		ctx.Abort500(syserrors.NewError("查询失败", err))
	}
	ctx.JSONOkH("查询成功！", H{
		"count": count,
	})
}

// @router /query [get]
func (ctx *MessageController) Query() {
	pageno, err := ctx.GetInt("pageno", 1)
	if err != nil || pageno < 1 {
		pageno = 1
	}
	limit, err := ctx.GetInt("limit", 10)
	if err != nil || limit < 5 {
		limit = 10
	}

	datas, err := ctx.Dao.QueryMessageForNoteByPage("", pageno, limit)
	if err != nil {
		ctx.Abort500(syserrors.NewError("查询失败", err))
	}
	ctx.JSONOkH("查询成功！", H{
		"data": datas,
	})
}
