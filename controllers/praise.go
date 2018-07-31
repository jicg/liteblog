package controllers

import (
	"github.com/jicg/liteblog/models"
	"github.com/jicg/liteblog/syserrors"
)

type PraiseController struct {
	BaseController
}

func (ctx *PraiseController) NestPrepare() {
	ctx.MustLogin()
}

// @router /:type/:key [post]
func (ctx *PraiseController) Parse() {
	key := ctx.Ctx.Input.Param(":key")
	ttype := ctx.Ctx.Input.Param(":type")
	var (
		praise  int = 0
		user_id int = int(ctx.User.ID)
		err     error
	)
	ctx.Dao.Begin()
	switch ttype {
	case "message":
		var message models.Message
		if message, err = ctx.Dao.QueryMessageByKey(key); err != nil {
			ctx.Dao.Rollback()
			ctx.Abort500(syserrors.NewError("点赞失败", err))
		}
		message.Praise = message.Praise + 1
		if err := ctx.Dao.UpdateMessage4Praise(&message); err != nil {
			ctx.Dao.Rollback()
			ctx.Abort500(syserrors.NewError("点赞失败", err))
		}
		praise = message.Praise
	case "note":
		var note models.Note
		if note, err = ctx.Dao.QueryNoteByKey(key); err != nil {
			ctx.Dao.Rollback()
			ctx.Abort500(syserrors.NewError("点赞失败", err))
		}
		note.Praise = note.Praise + 1
		if err := ctx.Dao.UpdateNote4Praise(&note); err != nil {
			ctx.Dao.Rollback()
			ctx.Abort500(syserrors.NewError("点赞失败", err))
		}
		praise = note.Praise
	default:
		ctx.Dao.Rollback()
		ctx.Abort500(syserrors.NewError("未知类型", nil))
	}

	p := models.PraiseLog{
		Key:    key,
		UserID: user_id,
		Type:   ttype,
	}
	var pp models.PraiseLog
	if pp, err = ctx.Dao.QueryPraiseLog(key, user_id, ttype); err != nil {
		pp = p
	} else {
		if pp.Flag {
			ctx.Dao.Rollback()
			ctx.Abort500(syserrors.HasPraiseError{})
		}
	}
	pp.Flag = true
	if err := ctx.Dao.SavePraiseLog(&pp); err != nil {
		ctx.Dao.Rollback()
		ctx.Abort500(syserrors.NewError("点赞失败", err))
	}
	ctx.Dao.Commit()
	ctx.JSONOkH("点赞成功！", H{
		"praise": praise,
	})
}
