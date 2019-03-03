package controller

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
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

// FindByName 名前でユーザを検索する
func (c *UserController) FindByName(w http.ResponseWriter, r *http.Request) {
	// inputPortの組み立て
	outputPort := c.OutputFactory(w)
	inputPort := c.InputFactory(outputPort)

	// クエリから名前を取得
	name := r.URL.Query().Get("name")

	// usecaseの実行
	inputPort.FindByName(name)
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
