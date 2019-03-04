package controller

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/grandcolline/clean-arch-demo/adapter/controller/form"
	"github.com/grandcolline/clean-arch-demo/adapter/logger"
	"github.com/grandcolline/clean-arch-demo/adapter/presenter"
	"github.com/grandcolline/clean-arch-demo/usecase"
)

// UserController ユーザコントローラ
type UserController struct {
	InputFactory  func(o usecase.UserOutputPort) usecase.UserInputPort
	OutputFactory func(w http.ResponseWriter) usecase.UserOutputPort
}

// NewUserController ユーザコントローラの作成
func NewUserController(repo usecase.UserRepositoryPort, logger logger.Logger) *UserController {
	return &UserController{
		InputFactory: func(o usecase.UserOutputPort) usecase.UserInputPort {
			return usecase.NewUserInteractor(o, repo, logger)
		},
		OutputFactory: func(w http.ResponseWriter) usecase.UserOutputPort {
			return presenter.NewUserPresenter(w)
		},
	}
}

// FindByID IDでユーザを検索する
func (c *UserController) FindByID(w http.ResponseWriter, r *http.Request) {
	// inputPortの組み立て
	outputPort := c.OutputFactory(w)
	inputPort := c.InputFactory(outputPort)

	// IDの取得
	userID := chi.URLParam(r, "userID")
	id, _ := strconv.ParseUint(userID, 10, 32)

	// usecaseの実行
	inputPort.FindByID(uint32(id))
}

// FindAll 全ユーザを検索する
func (c *UserController) FindAll(w http.ResponseWriter, r *http.Request) {
	// inputPortの組み立て
	outputPort := c.OutputFactory(w)
	inputPort := c.InputFactory(outputPort)

	// usecaseの実行
	inputPort.FindAll()
}

// Add 新規ユーザの追加
func (c *UserController) Add(w http.ResponseWriter, r *http.Request) {
	// inputPortの組み立て
	outputPort := c.OutputFactory(w)
	inputPort := c.InputFactory(outputPort)

	// POSTのデータを読み取る
	var f form.User
	if err := f.Set(r); err != nil {
		// TODO: 後でエラーハンドリングする
		return
	}

	// 必須・バリデーションチェック
	if ok, _ := f.Require(); !ok {
		// TODO: 後でエラーハンドリングする
		return
	}
	if ok, _ := f.Validate(); !ok {
		// TODO: 後でエラーハンドリングする
		return
	}

	// エンティティに詰める
	e := f.ToEntity()

	// usecaseの実行
	inputPort.Add(e)
}

// Change ユーザ情報の変更
func (c *UserController) Change(w http.ResponseWriter, r *http.Request) {
	// inputPortの組み立て
	outputPort := c.OutputFactory(w)
	inputPort := c.InputFactory(outputPort)

	// IDの取得
	userID := chi.URLParam(r, "userID")
	id, _ := strconv.ParseUint(userID, 10, 32)

	// POSTのデータを読み取る
	var f form.User
	if err := f.Set(r); err != nil {
		// TODO: 後でエラーハンドリングする
		return
	}

	// バリデーションチェック
	if ok, _ := f.Validate(); !ok {
		// TODO: 後でエラーハンドリングする
		return
	}

	// エンティティに詰める
	e := f.ToEntity()
	e.ID = uint32(id)

	// usecaseの実行
	inputPort.Change(e)
}

// Delete ユーザを削除する
func (c *UserController) Delete(w http.ResponseWriter, r *http.Request) {
	// inputPortの組み立て
	outputPort := c.OutputFactory(w)
	inputPort := c.InputFactory(outputPort)

	// IDの取得
	userID := chi.URLParam(r, "userID")
	id, _ := strconv.ParseUint(userID, 10, 32)

	// usecaseの実行
	inputPort.Delete(uint32(id))
}
