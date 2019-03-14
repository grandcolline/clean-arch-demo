package usecase

import "github.com/grandcolline/clean-arch-demo/entity"

// UserInteractor ユーザインタラクタ
type UserInteractor struct {
	UserOutputPort     UserOutputPort
	CmnOutputPort      CmnOutputPort
	UserRepositoryPort UserRepositoryPort
	Logger             Logger
}

// UserInputPort ユーザインプットポート
type UserInputPort interface {
	FindAll()
	FindByID(uint32)
	FindByName(string)
	Add(entity.User)
	Change(entity.User)
	Delete(uint32)
}

// UserOutputPort ユーザアウトプットポート
type UserOutputPort interface {
	RenderUser(*entity.User) error
	RenderUserList(*[]entity.User) error
}

// UserRepositoryPort ユーザレポジトリポート
type UserRepositoryPort interface {
	Store(entity.User) (entity.User, error)
	Update(entity.User) error
	FindAll() ([]entity.User, error)
	FindByName(string) ([]entity.User, error)
	FindByID(uint32) (entity.User, error)
	Delete(entity.User) error
	IsNotFound(error) bool
}

// NewUserInteractor はユーザインタラクタの作成を行います。
func NewUserInteractor(out UserOutputPort, cout CmnOutputPort, repo UserRepositoryPort, logger Logger) UserInputPort {
	return &UserInteractor{
		UserOutputPort:     out,
		CmnOutputPort:      cout,
		UserRepositoryPort: repo,
		Logger:             logger,
	}
}

// Add は新規ユーザを追加するメソッドです。
func (i *UserInteractor) Add(u entity.User) {
	i.Logger.Log("Interactor: User Add")

	// ユーザの登録
	user, err := i.UserRepositoryPort.Store(u)
	if err != nil {
		i.Logger.Log("error")
		i.CmnOutputPort.ServerErrRender()
		return
	}

	// Output
	i.UserOutputPort.RenderUser(&user)
}

/*
Change はユーザ情報の変更をします。

変更可能な値は名前とメールアドレスのみです。
*/
func (i *UserInteractor) Change(u entity.User) {
	i.Logger.Log("Interactor: User Change")

	// ユーザ情報の取得
	user, err := i.UserRepositoryPort.FindByID(u.ID)
	if i.UserRepositoryPort.IsNotFound(err) {
		i.Logger.Log("entity not found")
		i.CmnOutputPort.NoRecordErrRender("ユーザが存在しません。")
		return
	}
	if err != nil {
		i.Logger.Log("error")
		i.CmnOutputPort.ServerErrRender()
		return
	}

	// 情報の更新
	if u.Name != "" {
		user.Name = u.Name
	}

	if u.Email != "" {
		user.Email = u.Email
	}

	// ユーザ情報の変更
	if err := i.UserRepositoryPort.Update(user); err != nil {
		i.Logger.Log("error")
		return
	}

	// Output
	i.UserOutputPort.RenderUser(&user)
}

// FindAll はすべてのユーザを検索します。
func (i *UserInteractor) FindAll() {
	i.Logger.Log("Interactor: User FindAll")

	users, err := i.UserRepositoryPort.FindAll()
	if i.UserRepositoryPort.IsNotFound(err) {
		i.Logger.Log("entity not found")
		i.CmnOutputPort.NoRecordErrRender("ユーザが存在しません。")
		return
	}
	if err != nil {
		i.Logger.Log("error")
		i.CmnOutputPort.ServerErrRender()
		return
	}

	// Output
	i.UserOutputPort.RenderUserList(&users)
}

// FindByID はIDでユーザを検索します。
func (i *UserInteractor) FindByID(id uint32) {
	i.Logger.Log("Interactor: User FindByID")
	user, err := i.UserRepositoryPort.FindByID(id)
	if i.UserRepositoryPort.IsNotFound(err) {
		i.Logger.Log("entity not found")
		i.CmnOutputPort.NoRecordErrRender("ユーザが存在しません。")
		return
	}
	if err != nil {
		i.Logger.Log("error")
		i.CmnOutputPort.ServerErrRender()
		return
	}

	// Output
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

	// Output
	i.UserOutputPort.RenderUserList(&users)
}

// Delete はユーザを削除します。
func (i *UserInteractor) Delete(id uint32) {
	i.Logger.Log("Interactor: User Delete")
	// ユーザの取得
	user, err := i.UserRepositoryPort.FindByID(id)
	if i.UserRepositoryPort.IsNotFound(err) {
		i.Logger.Log("entity not found")
		i.CmnOutputPort.NoRecordErrRender("ユーザが存在しません。")
		return
	}
	if err != nil {
		i.Logger.Log("error")
		i.CmnOutputPort.ServerErrRender()
		return
	}

	// ユーザの削除
	if err = i.UserRepositoryPort.Delete(user); err != nil {
		i.Logger.Log("error")
		i.CmnOutputPort.ServerErrRender()
		return
	}

	// Output
	i.CmnOutputPort.SuccessRender()
}
