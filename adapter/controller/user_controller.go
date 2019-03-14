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
	InputFactory    func(out usecase.UserOutputPort, cout usecase.CmnOutputPort) usecase.UserInputPort
	CmnInputFactory func(w http.ResponseWriter) usecase.CmnOutputPort
	OutputFactory   func(w http.ResponseWriter) usecase.UserOutputPort
}

// NewUserController ユーザコントローラの作成
func NewUserController(repo usecase.UserRepositoryPort, logger logger.Logger) *UserController {
	return &UserController{
		InputFactory: func(out usecase.UserOutputPort, cout usecase.CmnOutputPort) usecase.UserInputPort {
			return usecase.NewUserInteractor(out, cout, repo, logger)
		},
		CmnInputFactory: func(w http.ResponseWriter) usecase.CmnOutputPort {
			return presenter.NewCmnPresenter(w)
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
	cmnOutputPort := c.CmnInputFactory(w)
	inputPort := c.InputFactory(outputPort, cmnOutputPort)

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
	cmnOutputPort := c.CmnInputFactory(w)
	inputPort := c.InputFactory(outputPort, cmnOutputPort)

	// usecaseの実行
	inputPort.FindAll()
}

// Add 新規ユーザの追加
func (c *UserController) Add(w http.ResponseWriter, r *http.Request) {
	// inputPortの組み立て
	outputPort := c.OutputFactory(w)
	cmnOutputPort := c.CmnInputFactory(w)
	inputPort := c.InputFactory(outputPort, cmnOutputPort)

	// POSTのデータを読み取る
	var f form.User
	if err := f.Set(r); err != nil {
		// TODO: 後でエラーハンドリングする
		return
	}

	// 必須チェック・バリデーションチェック
	if ok, messages := f.Require(); !ok {
		cmnOutputPort.ValidationErrRender(messages)
		return
	}
	if ok, messages := f.Validate(); !ok {
		cmnOutputPort.ValidationErrRender(messages)
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
	cmnOutputPort := c.CmnInputFactory(w)
	inputPort := c.InputFactory(outputPort, cmnOutputPort)

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
	if ok, messages := f.Validate(); !ok {
		cmnOutputPort.ValidationErrRender(messages)
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
	cmnOutputPort := c.CmnInputFactory(w)
	inputPort := c.InputFactory(outputPort, cmnOutputPort)

	// IDの取得
	userID := chi.URLParam(r, "userID")
	id, err := strconv.ParseUint(userID, 10, 32)
	if err != nil {
		cmnOutputPort.SuccessRender()
	}

	// usecaseの実行
	inputPort.Delete(uint32(id))
}
