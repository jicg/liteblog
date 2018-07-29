package models

import "github.com/jinzhu/gorm"

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
	Praise  int
	Content string
}

func QueryNoteByKeyAndUserId(key string, userid int) (*Note, error) {
	var note Note
	return &note, db.Model(&Note{}).Where("key = ? and user_id = ?", key, userid).Take(&note).Error
}

func QueryNoteByKey(key string) (*Note, error) {
	var note Note
	return &note, db.Model(&Note{}).Where("key = ? ", key).Take(&note).Error
}

func AllVisitCount(key string) error {
	return db.Model(&Note{}).Where("key = ?", key).Update("visit", gorm.Expr("visit + 1")).Error
}

func DelNoteByKey(key string, userid int) (error) {
	return db.Delete(Note{}, "key = ? and user_id = ? ", key, userid).Error
}
func QueryNotesBy(page, limit int) ([]*Note, error) {
	var note []*Note
	return note, db.Model(&Note{}).Offset(page * limit).Limit(limit).Order("updated_at DESC").Find(&note).Error
}

func SaveNote(n *Note) error {
	return db.Save(n).Error
}
