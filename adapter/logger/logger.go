package logger

import (
	"log"

	"github.com/grandcolline/clean-arch-demo/usecase"
)

// Logger ロガー
type Logger struct {
	level logLevel
}

// LogLevel ログレベル
type logLevel string

const (
	// debugLevel 開発に必要なデバッグ情報
	debugLevel logLevel = "DEBUG"
	// infoLevel 通常の情報
	infoLevel logLevel = "INFO"
	// errorLevel エラー時の情報
	errorLevel logLevel = "ERROR"
)

// NewLogger ロガーの作成
func NewLogger(s string) usecase.LoggerPort {
	lv := logLevel(s)
	// ログレベルが正しいものかどうかを確認
	if lv != debugLevel && lv != infoLevel && lv != errorLevel {
		panic("loglevel is invalid")
	}
	return &Logger{
		level: lv,
	}
}

// Debug はデバッグレベルのログを出力します。
func (l *Logger) Debug(args ...interface{}) {
	if l.outputCheck(debugLevel) {
		log.SetPrefix("[DEBUG] ")
		log.Println(args...)
	}
}

// Info はインフォレベルのログを出力します。
func (l *Logger) Info(args ...interface{}) {
	if l.outputCheck(infoLevel) {
		log.SetPrefix("[INFO] ")
		log.Println(args...)
	}
}

// Error はエラーレベルのログを出力します。
func (l *Logger) Error(args ...interface{}) {
	if l.outputCheck(errorLevel) {
		log.SetPrefix("[ERROR] ")
		log.SetFlags(log.LstdFlags | log.Lshortfile)
		log.Println(args...)
	}

}

// outputCheck は設定されているログレベルと比較して、出力するかどうかを判定します。
func (l *Logger) outputCheck(level logLevel) bool {
	switch l.level {
	// エラーレベルが設定されているとき
	case errorLevel:
		return level == errorLevel
	// インフォレベルが設定されているとき
	case infoLevel:
		return level == errorLevel || level == infoLevel
	// デバッグレベルが設定されているとき
	case debugLevel:
		return true
	default:
		return false
	}
}
