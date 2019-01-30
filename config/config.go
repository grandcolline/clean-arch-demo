package config

// FIXME: 標準パッケージだけで実装できないかなぁ？
import (
	"github.com/spf13/viper"
)

var config *viper.Viper
var defaults = map[string]string{
	// MySQL関連の設定
	"mysql_host":        "mysql",
	"mysql_slave_host":  "mysql",
	"mysql_db_name":     "app",
	"mysql_db_user":     "root",
	"mysql_db_password": "",
}

// Init is an exported method that takes the environment starts the viper
// (external lib) and returns the configuration struct.
func Init() {
	v := viper.New()
	for key, value := range defaults {
		v.SetDefault(key, value)
	}

	// APPが接頭辞の環境変数で上書き
	v.SetEnvPrefix("app")
	v.AutomaticEnv()
	config = v
}

// GetConfig は設定を取得します
func GetConfig() *viper.Viper {
	return config
}
