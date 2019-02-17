package usecase

import (
	"github.com/grandcolline/clean-arch-demo/entity"
	"github.com/grandcolline/clean-arch-demo/usecase/interfaces"
)

type UserInteractor struct {
	UserRepository interfaces.UserRepository
	Logger         interfaces.Logger
}

func (i *UserInteractor) Add(u entity.User) (int, error) {
	i.Logger.Log("store user!")
	return i.UserRepository.Store(u)
}

func (i *UserInteractor) FindByName(name string) ([]entity.User, error) {
	i.Logger.Log("find user!")
	return i.UserRepository.FindByName(name)
}
