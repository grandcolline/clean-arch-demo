package usecase

import (
	"github.com/grandcolline/clean-arch-demo/entity"
	"github.com/grandcolline/clean-arch-demo/usecase/inputport"
	"github.com/grandcolline/clean-arch-demo/usecase/outputport"
)

type userService struct {
	UserInputport  inputport.UserInputport
	UserOutputport outputport.UserOutputport
}

type UserService interface {
	Create(u *entity.User) error
	Get(u []*entity.User) ([]*entity.User, error)
}

func NewUserService(in inputport.UserInputport, out outputport.UserOutputport) UserService {
	return &userService{in, out}
}

func (userService *userService) Create(u *entity.User) error {
	return userService.UserInputport.Store(u)
}

func (userService *userService) Get(u []*entity.User) ([]*entity.User, error) {
	us, err := userService.UserInputport.FindAll(u)
	if err != nil {
		return nil, err
	}
	return userService.UserOutputport.ResponseUsers(us), nil
}
