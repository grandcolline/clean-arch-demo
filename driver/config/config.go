package config

import (
	"log"

	"github.com/kelseyhightower/envconfig"
)

// AppConf アプリケーション全体の設定
type AppConf struct {
	Port    string `default:"8080"`  // サーバ起動時に受け付けるポート
	Gateway string `default:"mysql"` // Gatewayのモード（mysql/mock）
}

// Init アプリケーション全体設定を環境変数から取得します
func (conf *AppConf) Init() {
	err := envconfig.Process("app", conf)
	if err != nil {
		log.Fatal(err.Error())
	}
}

// MysqlConf MySQLの設定
// これはAppConf.Gatewayがmysqlのときに読み込まれる。
type MysqlConf struct {
	Host        string `default:"mysql"` // 接続先ホスト
	Port        string `default:"3306"`  // 接続先ポート
	User        string `default:"root"`  // ユーザ
	Password    string `default:""`      // パスワード
	DBName      string `default:"app"`   // データベース名
	CreateTable bool   `default:"false"` // テーブルがないときにテーブルを作成するか（true/false）
}

// Init MySQLの設定を環境変数から取得します
func (conf *MysqlConf) Init() {
	err := envconfig.Process("mysql", conf)
	if err != nil {
		log.Fatal(err.Error())
	}
}
