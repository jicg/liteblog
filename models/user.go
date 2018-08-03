package models

//type Model struct {
//	ID        int        `gorm:"primary_key"`
//	CreatedAt time.Time
//	UpdatedAt time.Time
//	DeletedAt *time.Time `sql:"index"`
//}

//用户表
type User struct {
	Model
	Name   string `gorm:"unique_index" json:"name"`
	Email  string `gorm:"unique_index" json:"email"`
	Avatar string `json:"avatar"`
	Pwd    string `json:"-"`
	Role   int    `gorm:"default:0" json:"role"` // 0 管理员 1正常用户
}

func (db *DB) QueryUserByEmailAndPassword(email, password string) (user User, err error) {
	return user, db.db.Model(&User{}).Where("email = ? and pwd = ?", email, password).Take(&user).Error
}

func (db *DB) QueryUserByName(name string) (user User, err error) {
	return user, db.db.Where("name = ?", name).Take(&user).Error
}

func (db *DB) QueryUserByEmail(email string) (user User, err error) {
	return user, db.db.Where("email = ?", email).Take(&user).Error
}

func SaveUser(user *User) (error) {
	return db.Create(user).Error
}
