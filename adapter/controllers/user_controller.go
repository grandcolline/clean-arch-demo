package controllers

import (
	"net/http"

	"github.com/grandcolline/clean-arch-demo/adapter/gateway"
	"github.com/grandcolline/clean-arch-demo/adapter/interfaces"
	"github.com/grandcolline/clean-arch-demo/adapter/presenter"
	"github.com/grandcolline/clean-arch-demo/usecase"
	"github.com/jinzhu/gorm"
)

// UserController ユーザコントローラ
type UserController struct {
	input usecase.UserInputPort
}

// NewUserController ユーザコントローラの作成
func NewUserController(w http.ResponseWriter, conn *gorm.DB, logger interfaces.Logger) *UserController {
	// プレゼンタを作成
	out := &presenter.UserPresenter{
		Writer: w,
	}
	// レポジトリを作成
	repo := &gateway.UserRepository{
		Conn: conn,
	}
	return &UserController{
		input: usecase.NewUserInteractor(out, repo, logger),
	}
}

// FindByName 名前でユーザを検索する
func (c *UserController) FindByName(w http.ResponseWriter, r *http.Request) {
	// クエリから名前を取得
	name := r.URL.Query().Get("name")

	// usecaseの実行
	// 出力はコントローラは関与しない
	c.input.FindByName(name)
}
