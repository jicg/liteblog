package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var (
	db gorm.DB
)
func init() {
	db, err := gorm.Open("sqlite3", "data/data.db")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	// Migrate the schema
	db.AutoMigrate(&User{}, &Note{}, &Message{})

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
