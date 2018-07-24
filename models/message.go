package models

import "github.com/jinzhu/gorm"


type Note struct {
	gorm.Model
	UserID  int
	User    User
	Title   string
	Content string
}

type Message struct {
	gorm.Model
	UserID  int
	User    User
	NoteID  int
	Note    Note
	Praise  int
	Content string
}
