package models

import (
	"time"
)


type Model struct {
	ID        int        `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
}

type User struct {
	Model
	Name   string `gorm:"unique_index"`
	Email  string `gorm:"unique_index"`
	Avatar string
	Pwd    string
	Role   int    `gorm:"default:0"` // 0 管理员 1正常用户
}

type Message struct {
	Model
	UserID  int
	User    User
	NoteID  int
	Note    Note
	Praise  int
	Content string
}

type Note struct {
	Model
	UserID  int
	User    User
	Title   string
	Content string
}
