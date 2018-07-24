package controllers

import (
	"github.com/jicg/liteblog/models"
	"github.com/jicg/liteblog/syserrors"
	"strings"
	"errors"
)

type UserController struct {
	BaseController
}

// @router /login [post]
func (c *UserController) Login() {
	email := c.GetMustString("email", "邮箱不能为空！")
	pwd := c.GetMustString("password", "密码不能为空！")
	var (
		user *models.User
		err  error
	)
	if user, err = models.QueryUserByEmailAndPassword(email, pwd); err != nil {
		c.Abort500(syserrors.NewError("邮箱或密码不对", err))
	}
	c.SetSession(SESSION_USER_KEY, user)
	c.JSONOk("登陆成功", "/")
}

// @router /reg [post]
func (c *UserController) Reg() {
	name := c.GetMustString("name", "昵称不能为空！")
	email := c.GetMustString("email", "邮箱不能为空！")
	pwd1 := c.GetMustString("password", "密码不能为空！")
	pwd2 := c.GetMustString("password2", "确认密码不能为空！")
	if strings.Compare(pwd1, pwd2) != 0 {
		c.Abort500(errors.New("密码与确认密码 必须要一致！"))
	}
	if u, err := models.QueryUserByName(name); err == nil && u != nil && u.ID != 0 {
		c.Abort500(syserrors.NewError("用户昵称已经存在!", err))
	}
	if u, err := models.QueryUserByEmail(email); err == nil && u != nil && u.ID != 0 {
		c.Abort500(syserrors.NewError("用户邮箱已经存在！", err))
	}

	if err := models.SaveUser(&models.User{
		Name:   name,
		Email:  email,
		Pwd:    pwd1,
		Avatar: "",
		Role:   1,
	}); err != nil {
		c.Abort500(syserrors.NewError("用户注册失败", err))
	}
	c.JSONOk("注册成功", "/user")
}

// @router /logout [get]
func (c *UserController) Logout() {
	c.MustLogin()
	c.DelSession(SESSION_USER_KEY)
	c.Redirect("/", 302)
}
