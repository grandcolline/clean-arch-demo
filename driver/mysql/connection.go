package mysql

import (
	"github.com/grandcolline/clean-arch-demo/adapter/gateway"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

// Connect MySQLに接続する
func Connect() *gorm.DB {
	var err error

	db, err = gorm.Open("mysql", "root:@tcp(mysql:3306)/app?parseTime=true")
	if err != nil {
		panic(err)
	}
	// テーブルの存在チェック
	if !db.HasTable(&gateway.User{}) {
		if err := db.Table("users").CreateTable(&gateway.User{}).Error; err != nil {
			panic(err)
		}
	}
	return db
}

// CloseConn MySQLとの接続を切断する
func CloseConn() {
	db.Close()
}
