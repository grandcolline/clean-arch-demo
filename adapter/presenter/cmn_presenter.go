package presenter

import (
	"encoding/json"
	"net/http"

	"github.com/grandcolline/clean-arch-demo/usecase"
)

// CmnPresenter 共通プレゼンタ
type CmnPresenter struct {
	writer http.ResponseWriter
}

// NewCmnPresenter は共通プレゼンタを作成します。
func NewCmnPresenter(w http.ResponseWriter) usecase.CmnOutputPort {
	return &CmnPresenter{
		writer: w,
	}
}

// Cmn 共通レスポンス
type Cmn struct {
	Result string  `json:"result"`
	Status int     `json:"status"`
	Errors []Error `json:"errors"`
}

// Error エラー内容レスポンス
type Error struct {
	// Code    int    `json:"code"`
	Message string `json:"message"`
}

/*
SuccessRender は処理の成功時のレスポンスを返します。

Example Response:
	{
		"result": 200
		"status": "success"
	}
*/
func (p *CmnPresenter) SuccessRender() {
	statusCode := http.StatusOK //200
	cmn := &Cmn{
		Status: statusCode,
		Result: "success",
	}
	res, err := json.Marshal(cmn)
	if err != nil {
		http.Error(p.writer, err.Error(), http.StatusInternalServerError)
		return
	}

	p.writer.Header().Set("Content-Type", "application/json")
	p.writer.WriteHeader(statusCode)
	p.writer.Write(res)
}

/*
ValidationErrRender はバリデーションエラーを返します。

Example Response:
	{
		"result": 422
		"status": "validation error"
		"errors": [
			{ "message": "ユーザ名は必須項目です。" },
			{ "message": "メールアドレスは必須項目です。" }
		]
	}
*/
func (p *CmnPresenter) ValidationErrRender(ss []string) {
	statusCode := http.StatusUnprocessableEntity // 422
	var errors []Error
	for _, s := range ss {
		errors = append(
			errors,
			Error{
				Message: s,
			},
		)
	}
	cmn := &Cmn{
		Status: statusCode,
		Result: "validation error",
		Errors: errors,
	}
	res, err := json.Marshal(cmn)
	if err != nil {
		http.Error(p.writer, err.Error(), http.StatusInternalServerError)
		return
	}
	p.writer.Header().Set("Content-Type", "application/json")
	p.writer.WriteHeader(statusCode)
	p.writer.Write(res)
}

/*
ServerErrRender はサーバエラーを返します。

Example Response:
	{
		"result": 503
		"status": "internal server error"
	}
*/
func (p *CmnPresenter) ServerErrRender() {
	statusCode := http.StatusServiceUnavailable // 503
	cmn := &Cmn{
		Status: statusCode,
		Result: "internal server error",
	}
	res, err := json.Marshal(cmn)
	if err != nil {
		http.Error(p.writer, err.Error(), http.StatusInternalServerError)
		return
	}
	p.writer.Header().Set("Content-Type", "application/json")
	p.writer.WriteHeader(statusCode)
	p.writer.Write(res)
}

/*
NoRecordErrRender はレコードがない場合に返します。

Example Response:
	{
		"result": 404
		"status": "no record error"
		"errors": [
			{ "message": "ユーザが存在しません。" }
		]
	}
*/
func (p *CmnPresenter) NoRecordErrRender(s string) {
	statusCode := http.StatusNotFound // 404
	var errors []Error
	errors = append(
		errors,
		Error{
			Message: s,
		},
	)
	cmn := &Cmn{
		Status: statusCode,
		Result: "no record error",
		Errors: errors,
	}
	res, err := json.Marshal(cmn)
	if err != nil {
		http.Error(p.writer, err.Error(), http.StatusInternalServerError)
		return
	}
	p.writer.Header().Set("Content-Type", "application/json")
	p.writer.WriteHeader(statusCode)
	p.writer.Write(res)
}
