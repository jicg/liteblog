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
	if err := db.Model(&Note{}).Where("key = ? and user_id = ?", key, userid).Take(&note).Error; err != nil {
		return nil, err
	}
	return &note, nil
}

func QueryNoteByKey(key string) (*Note, error) {
	var note Note
	if err := db.Model(&Note{}).Where("key = ? ", key).Take(&note).Error; err != nil {
		return nil, err
	}
	return &note, nil
}

func QueryNotesBy(page, limit int) ([]*Note, error) {
	var note []*Note
	return note,db.Model(&Note{}).Offset(page * limit).Limit(limit).Order("updated_at DESC").Find(&note).Error
	//if err := ; err != nil {
	//	return nil, err
	//}
	//return note, nil
}

func SaveNote(n *Note) error {
	return db.Save(n).Error
}
