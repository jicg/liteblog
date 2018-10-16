package models

type PraiseLog struct {
	Model
	Key    string `sql:"index"`
	UserID int    `sql:"index"`
	Type   string `sql:"index"`
	Flag   bool
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
	return parselog, db.db.Where("`key` = ? and user_id =? and type = ? ", key, user_id, ttype).Take(&parselog).Error
}

func (db *DB) SavePraiseLog(p *PraiseLog) error {
	return db.db.Save(&p).Error
}
