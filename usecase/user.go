package usecase

import (
	"github.com/grandcolline/clean-arch-demo/entity"
)

// UserInteractor ユーザインタラクタ
type UserInteractor struct {
	UserRepository UserRepository
	Logger         Logger
}

// UserRepository ユーザレポジトリ
type UserRepository interface {
	Store(entity.User) (int, error)
	FindByName(string) ([]entity.User, error)
	FindAll() ([]entity.User, error)
}

// Add 新規ユーザを追加する
func (i *UserInteractor) Add(u entity.User) (int, error) {
	i.Logger.Log("store user!")
	return i.UserRepository.Store(u)
}

// FindByName 名前でユーザを検索する
func (i *UserInteractor) FindByName(name string) ([]entity.User, error) {
	i.Logger.Log("find user!")
	return i.UserRepository.FindByName(name)
}
