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

type Message struct {
	Model
	Key    string `gorm:"unique_index;not null;" json:"key"`
	UserID int    `json:"user_id"`
	User   User   `json:"user"`
	//NoteID  int
	//Note    Note
	NoteKey string `sql:"index" json:"note_key"`
	Content string `json:"content"`
	Praise  int    `gorm:"default:0" json:"praise"`
}

type PraiseLog struct {
	Model
	Key    string `sql:"index"`
	UserID int    `sql:"index"`
	Type   string `sql:"index"`
	Flag   bool
}

func (db *DB) QueryNoteByKeyAndUserId(key string, userid int) (note Note, err error) {
	return note, db.db.Model(&Note{}).Where("key = ? and user_id = ?", key, userid).Take(&note).Error
}

func (db *DB) QueryNoteByKey(key string) (note Note, err error) {
	return note, db.db.Model(&Note{}).Where("key = ? ", key).Take(&note).Error
}

func (db *DB) AllVisitCount(key string) error {
	return db.db.Model(&Note{}).Where("key = ?", key).UpdateColumn("visit", gorm.Expr("visit + 1")).Error
}

func (db *DB) DelNoteByKey(key string, userid int) (error) {
	return db.db.Delete(Note{}, "key = ? and user_id = ? ", key, userid).Error
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

func (db *DB) UpdateNote4Praise(n *Note) error {
	return db.db.Model(&Note{}).UpdateColumn("praise", n.Praise).Error
}

//
//func (dbt *DB) PraiseOK(p *PraiseLog) (int, error) {
//	db = dbt.db.Begin()
//	var praise int = 0
//	switch p.Type {
//	case "message":
//		var message Message
//		if err := db.Where("key = ? ", p.Key).Take(&message).Error; err != nil {
//			db.Callback()
//			return praise, err
//		}
//		message.Praise = message.Praise + 1
//		if err := db.Save(message).Error; err != nil {
//			db.Callback()
//			return praise, err
//		}
//		praise = message.Praise
//	case "note":
//		var note Note
//		if err := db.Where("key = ?", p.Key).Take(&note).Error; err != nil {
//			db.Callback()
//			return praise, err
//		}
//		note.Praise = note.Praise + 1
//		if err := db.Save(&note).Error; err != nil {
//			db.Callback()
//			return praise, err
//		}
//		praise = note.Praise
//	default:
//		db.Callback()
//		return 0, errors.New("未知类型")
//	}
//	var pp PraiseLog
//	if db.Where("key = ? and user_id =? and type = ? ", p.Key, p.UserID, p.Type).Take(&pp).RecordNotFound() {
//		pp = *p
//	} else {
//		if pp.Flag {
//			db.Callback()
//			return 0, errors.New("您已经点过赞！")
//		}
//	}
//	pp.Flag = true
//	if err := db.Save(&pp).Error; err != nil {
//		db.Callback()
//		return praise, err
//	}
//	db.Commit()
//	return praise, nil
//}

func (db *DB) QueryPraiseLog(key string, user_id int, ttype string) (parselog PraiseLog, err error) {
	return parselog, db.db.Where("key = ? and user_id =? and type = ? ", key, user_id, ttype).Take(&parselog).Error
}

func (db *DB) SavePraiseLog(p *PraiseLog) error {
	return db.db.Save(&p).Error
}

func (db *DB) QueryMessageByKey(key string) (message Message, err error) {
	return message, db.db.Model(&Message{}).Where("key = ? ", key).Take(&message).Error
}

//func (db *DB) QueryMessageForAdmin() (messages []*Message, err error) {
//	return messages, db.db.Where("note_key is null").Find(&messages).Error
//}

func (db *DB) QueryMessageForNote(key string) (messages []*Message, err error) {
	return messages, db.db.Preload("User").Where("note_key = ?", key).Order("updated_at desc").Find(&messages).Error
}

func (db *DB) SaveMessage(n *Message) error {
	return db.db.Save(n).Error
}

func (db *DB) UpdateMessage4Praise(n *Message) error {
	return db.db.Model(&Message{}).UpdateColumn("praise", n.Praise).Error
}
