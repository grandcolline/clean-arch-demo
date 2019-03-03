package controller

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/grandcolline/clean-arch-demo/adapter/logger"
	"github.com/grandcolline/clean-arch-demo/adapter/presenter"
	"github.com/grandcolline/clean-arch-demo/entity"
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

// Add 新規ユーザの追加
func (c *UserController) Add(w http.ResponseWriter, r *http.Request) {
	// inputPortの組み立て
	outputPort := c.OutputFactory(w)
	inputPort := c.InputFactory(outputPort)

	// POSTのデータからエンティティの作成
	var user entity.User
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		//TODO: エラーハンドリングは後で考える
		panic(err)
	}
	if err := r.Body.Close(); err != nil {
		//TODO: エラーハンドリングは後で考える
		panic(err)
	}
	if err := json.Unmarshal(body, &user); err != nil {
		// w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		// w.WriteHeader(422) // unprocessable entity
		// if err := json.NewEncoder(w).Encode(err); err != nil {
		// 	panic(err)
		// }
		//TODO: エラーハンドリングは後で考える
		panic(err)
	}

	// usecaseの実行
	inputPort.Add(user)
}
