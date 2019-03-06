package mysql

import (
	"fmt"

	"github.com/grandcolline/clean-arch-demo/adapter/gateway/model"
	"github.com/grandcolline/clean-arch-demo/driver/config"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

// conf MySQLの接続設定
var conf config.MysqlConf

// Connect MySQLに接続する
func Connect() (db *gorm.DB) {
	var err error

	// DB設定の読み込み
	conf.Init()

	// DB接続
	url := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		conf.User,
		conf.Password,
		conf.Host,
		conf.Port,
		conf.DBName,
	)
	if db, err = gorm.Open("mysql", url); err != nil {
		panic(err)
	}

	// テーブルの存在チェック
	// TODO: db.AutoMigrateの方が良さそう...
	checkTable(db, "users", &model.User{})

	return
}

// CloseConn MySQLとの接続を切断する
func CloseConn() {
	db.Close()
}

// checkTable テーブルが存在するかチェックし、作成を行います。
func checkTable(db *gorm.DB, table string, model interface{}) {
	if !db.HasTable(model) {
		// conf.CreateTableをチェックし、テーブルを作成する。
		if conf.CreateTable {
			if err := db.Table(table).CreateTable(model).Error; err != nil {
				panic(err)
			}
		} else {
			panic(table + "テーブルが存在しません")
		}
	}
}
