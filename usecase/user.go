package usecase

import "github.com/grandcolline/clean-arch-demo/entity"

// UserInteractor ユーザインタラクタ
type UserInteractor struct {
	UserOutputPort     UserOutputPort
	UserRepositoryPort UserRepositoryPort
	Logger             Logger
}

// UserInputPort ユーザインプットポート
// usecaseの入力ポート。adpter層のcontrollerで使われる。
type UserInputPort interface {
	FindByName(string)
}

// UserOutputPort ユーザアウトプットポート
// usecaseの出力ポート。実装はadpter層のpresenter。
type UserOutputPort interface {
	RenderUser(*entity.User) error
	RenderUserList(*[]entity.User) error
}

// UserRepositoryPort ユーザレポジトリポート
// データストアとの接続で用いるポート。実装はadpter層のgateway。
type UserRepositoryPort interface {
	FindByName(string) ([]entity.User, error)
	// Store(entity.User) (int, error)
	// FindAll() ([]entity.User, error)
}

// NewUserInteractor ユーザインタラクタの作成
// ここでは入力ポートを作成している
func NewUserInteractor(out UserOutputPort, repo UserRepositoryPort, logger Logger) UserInputPort {
	return &UserInteractor{
		UserOutputPort:     out,
		UserRepositoryPort: repo,
		Logger:             logger,
	}
}

// Add 新規ユーザを追加する
// func (i *UserInteractor) Add(u entity.User) (int, error) {
// 	i.Logger.Log("store user!")
// 	return i.UserRepositoryPort.Store(u)
// }

// FindByName 名前でユーザを検索する
func (i *UserInteractor) FindByName(name string) {
	i.Logger.Log("find user by name!")
	users, err := i.UserRepositoryPort.FindByName(name)
	if err != nil {
		i.Logger.Log("error")
		return
	}
	user := users[0]
	i.UserOutputPort.RenderUser(&user)
}
