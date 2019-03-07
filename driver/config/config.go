package config

import (
	"log"

	"github.com/kelseyhightower/envconfig"
)

// AppConf アプリケーション全体の設定
type AppConf struct {
	Port    string `default:"8080"` // サーバ起動時に受け付けるポート
	Gateway string `default:"mock"` // Gatewayのモード（mock/mysql）
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
	Host        string `required:"true"` // 接続先ホスト
	Port        string `default:"3306"`  // 接続先ポート
	User        string `required:"true"` // DB接続ユーザ
	Password    string `required:"true"` // DB接続パスワード
	Database    string `required:"true"` // データベース名
	CreateTable bool   `default:"false"` // テーブルがないときにテーブルを作成するか（true/false）
}

// Init MySQLの設定を環境変数から取得します
func (conf *MysqlConf) Init() {
	err := envconfig.Process("mysql", conf)
	if err != nil {
		log.Fatal(err.Error())
	}
}
