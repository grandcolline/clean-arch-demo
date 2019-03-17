package config

import (
	"log"

	"github.com/kelseyhightower/envconfig"
)

// AppConf アプリケーション全体の設定
type AppConf struct {
	Port  string `default:"8080"`  // サーバ起動時に受け付けるポート
	Debug bool   `default:"false"` // デバッグモードで起動するか
}

// Init アプリケーション全体設定を環境変数から取得します
func (conf *AppConf) Init() {
	err := envconfig.Process("app", conf)
	if err != nil {
		log.Fatal(err.Error())
	}
}

// MysqlConf MySQLの設定
type MysqlConf struct {
	Host        string `required:"true"` // 接続先ホスト
	Port        string `default:"3306"`  // 接続先ポート
	User        string `required:"true"` // DB接続ユーザ
	Password    string `required:"true"` // DB接続パスワード
	Database    string `required:"true"` // データベース名
	LogMode     bool   `default:"false"` // SQLログを出力するか
	CreateTable bool   `default:"false"` // テーブルがないときにテーブルを作成するか
}

// Init MySQLの設定を環境変数から取得します
func (conf *MysqlConf) Init() {
	err := envconfig.Process("mysql", conf)
	if err != nil {
		log.Fatal(err.Error())
	}
}
