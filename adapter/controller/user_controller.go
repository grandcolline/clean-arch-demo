package controller

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/grandcolline/clean-arch-demo/adapter/gateway"
	"github.com/grandcolline/clean-arch-demo/adapter/logger"
	"github.com/grandcolline/clean-arch-demo/adapter/presenter"
	"github.com/grandcolline/clean-arch-demo/usecase"
	"github.com/jinzhu/gorm"
)

// UserController ユーザコントローラ
type UserController struct {
	input usecase.UserInputPort
}

// NewUserController ユーザコントローラの作成
func NewUserController(w http.ResponseWriter, conn *gorm.DB, logger logger.Logger) *UserController {
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

// FindByID IDでユーザを検索する
func (c *UserController) FindByID(w http.ResponseWriter, r *http.Request) {
	userID := chi.URLParam(r, "userID")
	id, _ := strconv.ParseUint(userID, 10, 32)

	// usecaseの実行
	// 出力はコントローラは関与しない
	c.input.FindByID(uint32(id))
}

// FindAll 全ユーザを検索する
func (c *UserController) FindAll(w http.ResponseWriter, r *http.Request) {
	c.input.FindAll()
}
