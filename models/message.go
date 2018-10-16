package models

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

func (db *DB) QueryMessageByKey(key string) (message Message, err error) {
	return message, db.db.Model(&Message{}).Where("`key` = ? ", key).Take(&message).Error
}

//func (db *DB) QueryMessageForAdmin() (messages []*Message, err error) {
//	return messages, db.db.Where("note_key is null").Find(&messages).Error
//}

func (db *DB) QueryMessageForNote(key string) (messages []*Message, err error) {
	return messages, db.db.Preload("User").Where("note_key = ?", key).Order("updated_at desc").Find(&messages).Error
}

func (db *DB) QueryMessageForNoteByPage(key string, page, limit int) (messages []*Message, err error) {
	return messages, db.db.Preload("User").Where("note_key = ?", key).Offset((page - 1) * limit).Limit(limit).Order("updated_at desc").Find(&messages).Error
}

func (db *DB) QueryMessageForNoteCount(key string) (count int, err error) {
	return count, db.db.Model(&Message{}).Where("note_key = ?", key).Count(&count).Error
}

func (db *DB) SaveMessage(n *Message) error {
	return db.db.Save(n).Error
}

func (db *DB) UpdateMessage4Praise(n *Message) error {
	return db.db.Model(&Message{}).UpdateColumn("praise", n.Praise).Error
}
