package models

import (
	"github.com/jinzhu/gorm"
)

//type Model struct {
//	ID        int        `gorm:"primary_key"`
//	CreatedAt time.Time
//	UpdatedAt time.Time
//	DeletedAt *time.Time `sql:"index"`
//}

//用户表
type User struct {
	gorm.Model
	Name   string `gorm:"unique_index"`
	Email  string `gorm:"unique_index"`
	Avatar string
	Pwd    string
	Role   int    `gorm:"default:0"` // 0 管理员 1正常用户
}

func QueryUserByEmailAndPassword(email, password string) (*User, error) {
	var user User
	if err := db.Model(&User{}).Where("email = ? and pwd = ?", email, password).Take(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func QueryUserByName(name string) (*User, error) {
	var user User
	if err := db.Where("name = ?", name).Take(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func QueryUserByEmail(email string) (*User, error) {
	var user User
	if err := db.Where("email = ?", email).Take(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func SaveUser(user *User) (error) {
	return db.Create(user).Error
}
