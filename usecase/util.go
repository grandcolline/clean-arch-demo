package usecase

/*
CmnOutputPort 共通アウトプットポート

複数のインタラクたで使うための、共通のアウトプットポート。
*/
type CmnOutputPort interface {
	SuccessRender()
	ValidationErrRender([]string)
	ServerErrRender()
	NoRecordErrRender(string)
}

// LoggerPort ログ出力ポート
type LoggerPort interface {
	Debug(args ...interface{})
	Info(args ...interface{})
	Error(args ...interface{})
}
