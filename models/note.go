package models

import (
	"github.com/jinzhu/gorm"
	"fmt"
)

type Note struct {
	Model
	Key     string `gorm:"unique_index;not null;"`
	UserID  int
	User    User
	Title   string
	Summary string `gorm:"type:text"`
	Content string `gorm:"type:text"`
	Files   string `gorm:"type:text"`
	Visit   int    `gorm:"default:0"`
	Praise  int    `gorm:"default:0"`
}

func (db *DB) QueryNoteByKeyAndUserId(key string, userid int) (note Note, err error) {
	return note, db.db.Model(&Note{}).Where("`key` = ? and user_id = ?", key, userid).Take(&note).Error
}

func (db *DB) QueryNoteByKey(key string) (note Note, err error) {
	return note, db.db.Model(&Note{}).Where("`key` = ? ", key).Take(&note).Error
}

func (db *DB) AllVisitCount(key string) error {
	return db.db.Model(&Note{}).Where("`key` = ?", key).UpdateColumn("visit", gorm.Expr("visit + 1")).Error
}

func (db *DB) DelNoteByKey(key string, userid int) (error) {
	return db.db.Delete(Note{}, "`key` = ? and user_id = ? ", key, userid).Error
}
func (db *DB) QueryNotesByPage(page, limit int, title string) (note []*Note, err error) {
	return note, db.db.Model(&Note{}).Where("title like ?", fmt.Sprintf("%%%s%%", title)).Offset((page - 1) * limit).Limit(limit).Order("updated_at DESC").Find(&note).Error
}
func (db *DB) QueryNotesCount(title string) (cnt int, err error) {
	return cnt, db.db.Model(&Note{}).Where("title like ?", fmt.Sprintf("%%%s%%", title)).Offset(-1).Limit(-1).Count(&cnt).Error
}

func (db *DB) SaveNote(n *Note) error {
	return db.db.Save(n).Error
}
