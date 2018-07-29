package models

import (
	"github.com/jinzhu/gorm"
	"fmt"
)

type Note struct {
	gorm.Model
	Key     string `gorm:"unique_index;not null;"`
	UserID  int
	User    User
	Title   string
	Summary string `gorm:"type:text"`
	Content string `gorm:"type:text"`
	Visit   int    `gorm:"default:0"`
	Praise  int    `gorm:"default:0"`
}

type Message struct {
	gorm.Model
	UserID  int
	User    User
	NoteID  int
	Note    Note
	NoteKey string `sql:"index"`
	Content string
	Praise  int    `gorm:"default:0"`
}

func QueryNoteByKeyAndUserId(key string, userid int) (note Note, err error) {
	return note, db.Model(&Note{}).Where("key = ? and user_id = ?", key, userid).Take(&note).Error
}

func QueryNoteByKey(key string) (note Note, err error) {
	return note, db.Model(&Note{}).Where("key = ? ", key).Take(&note).Error
}

func AllVisitCount(key string) error {
	return db.Model(&Note{}).Where("key = ?", key).Update("visit", gorm.Expr("visit + 1")).Error
}

func DelNoteByKey(key string, userid int) (error) {
	return db.Delete(Note{}, "key = ? and user_id = ? ", key, userid).Error
}
func QueryNotesByPage(page, limit int, title string) (note []*Note, err error) {
	return note, db.Model(&Note{}).Where("title like ?", fmt.Sprintf("%%%s%%", title)).Offset((page - 1)*limit).Limit(limit).Order("updated_at DESC").Find(&note).Error
}
func QueryNotesCount(title string) (cnt int, err error) {
	return cnt, db.Model(&Note{}).Where("title like ?", fmt.Sprintf("%%%s%%", title)).Offset(-1).Limit(-1).Count(&cnt).Error
}

func SaveNote(n *Note) error {
	return db.Save(n).Error
}
