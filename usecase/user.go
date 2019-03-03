package usecase

import "github.com/grandcolline/clean-arch-demo/entity"

// UserInteractor ユーザインタラクタ
type UserInteractor struct {
	UserOutputPort     UserOutputPort
	UserRepositoryPort UserRepositoryPort
	Logger             Logger
}

// UserInputPort ユーザインプットポート
//
// 入力ポートのInterface。
// adpter層のcontrollerで使われる。
type UserInputPort interface {
	FindAll()
	FindByID(uint32)
	FindByName(string)
	Add(entity.User)
}

// UserOutputPort ユーザアウトプットポート
//
// 出力ポートのInterface。
// 実装はadpter層のpresenter。
type UserOutputPort interface {
	RenderUser(*entity.User) error
	RenderUserList(*[]entity.User) error
}

// UserRepositoryPort ユーザレポジトリポート
//
// データストアとの接続で用いるポートのInterface。
// 実装はadpter層のgateway。
type UserRepositoryPort interface {
	Store(entity.User) (uint32, error)
	FindAll() ([]entity.User, error)
	FindByName(string) ([]entity.User, error)
	FindByID(uint32) (entity.User, error)
}

// NewUserInteractor ユーザインタラクタの作成
func NewUserInteractor(out UserOutputPort, repo UserRepositoryPort, logger Logger) UserInputPort {
	return &UserInteractor{
		UserOutputPort:     out,
		UserRepositoryPort: repo,
		Logger:             logger,
	}
}

// Add 新規ユーザを追加する
func (i *UserInteractor) Add(u entity.User) {
	i.Logger.Log("Interactor: User Add")

	// ユーザの登録
	id, err := i.UserRepositoryPort.Store(u)
	if err != nil {
		i.Logger.Log("error")
		return
	}

	// 作成したユーザの取得
	user, err := i.UserRepositoryPort.FindByID(id)
	if err != nil {
		i.Logger.Log("error")
		return
	}

	// Output
	i.UserOutputPort.RenderUser(&user)
}

// FindAll ずべてのユーザを検索する
func (i *UserInteractor) FindAll() {
	i.Logger.Log("Interactor: User FindAll")

	users, err := i.UserRepositoryPort.FindAll()
	if err != nil {
		i.Logger.Log("error")
		return
	}
	i.UserOutputPort.RenderUserList(&users)
}

// FindByID IDでユーザを検索する
func (i *UserInteractor) FindByID(id uint32) {
	i.Logger.Log("Interactor: User FindByID")
	user, err := i.UserRepositoryPort.FindByID(id)
	if err != nil {
		i.Logger.Log("error")
		return
	}
	i.UserOutputPort.RenderUser(&user)
}

// FindByName 名前でユーザを検索する
func (i *UserInteractor) FindByName(name string) {
	i.Logger.Log("Interactor: User FindByName")
	users, err := i.UserRepositoryPort.FindByName(name)
	if err != nil {
		i.Logger.Log("error")
		return
	}
	i.UserOutputPort.RenderUserList(&users)
}
