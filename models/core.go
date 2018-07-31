package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"os"
	"github.com/astaxie/beego/logs"
	"time"
)

type DB struct {
	db *gorm.DB
}

func (db *DB) Begin() {
	db.db = db.db.Begin()
}

func (db *DB) Rollback() {
	db.db = db.db.Rollback()
}

func (db *DB) Commit() {
	db.db = db.db.Commit()
}

var (
	db *gorm.DB
)

func NewDB() *DB {
	return &DB{db: db}
}

func init() {
	var err error
	// 创建data目录
	if err = os.MkdirAll("data", 0777); err != nil {
		panic("failed to connect database," + err.Error())
	}
	db, err = gorm.Open("sqlite3", "data/data.db")
	if err != nil {
		panic("failed to connect database")
	}
	// 自动同步表结构
	db.SetLogger(logs.GetLogger("orm"))
	db.LogMode(true)
	db.AutoMigrate(&User{}, &Note{}, &Message{}, &PraiseLog{})
	// Model(&User{})查询用户表, Count(&count) 将用户表的数据赋值给count字段。
	var count int
	if err := db.Model(&User{}).Count(&count).Error; err == nil && count == 0 {
		db.Create(&User{Name: "admin",
			//邮箱
			Email: "admin@qq.com",
			//密码
			Pwd: "123123",
			//头像地址
			Avatar: "/static/images/info-img.png",
			//是否认证 例： lyblog 作者
			Role: 0,
		})
	}
}

type Model struct {
	ID        uint       `gorm:"primary_key" json:"id"`
	CreatedAt time.Time  `json:"createtime"`
	UpdatedAt time.Time  `json:"updatetime"`
	DeletedAt *time.Time `sql:"index" json:"-"`
}
