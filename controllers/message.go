package controllers

import (
	"github.com/go-xweb/uuid"
	"github.com/jicg/liteblog/models"
	"github.com/jicg/liteblog/syserrors"
)

type MessageController struct {
	BaseController
}

func (ctx *MessageController) NestPrepare() {
	ctx.MustLogin()
}

// @router /new/?:key [post]
func (ctx *MessageController) NewMessage() {
	ctx.MustLogin()
	key := uuid.NewUUID().String()
	content := ctx.GetMustString("content", "内容不能为空")
	notekey := ctx.Ctx.Input.Param(":key")
	m := &models.Message{
		UserID:  int(ctx.User.ID),
		User:    *ctx.User,
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
