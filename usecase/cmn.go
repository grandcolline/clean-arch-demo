package usecase

// CmnOutputPort 共通アウトプットポート
type CmnOutputPort interface {
	SuccessRender()
	ValidationErrRender([]string)
	ServerErrRender()
	NoRecordErrRender(string)
}
