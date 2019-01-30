package datastore

import (
	"fmt"

	"github.com/grandcolline/clean-arch-demo/config"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func NewMySqlDB() *gorm.DB {

	url := fmt.Sprintf(
		"%s:%s@tcp(%s:3306)/%s?charset=utf8&parseTime=True&loc=Local",
		config.GetConfig().Get("mysql_db_user"),
		config.GetConfig().Get("mysql_db_password"),
		config.GetConfig().Get("mysql_host"),
		config.GetConfig().Get("mysql_db_name"),
	)

	conn, err := gorm.Open("mysql", url)

	if nil != err {
		panic(err)
	}

	// 応答確認
	err = conn.DB().Ping()
	if nil != err {
		panic(err)
	}
	// sqlログの詳細を出力
	conn.LogMode(true)

	// DBのエンジンを設定
	conn.Set("gorm:table_options", "ENGINE=InnoDB")

	return conn
}
