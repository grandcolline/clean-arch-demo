package controller

import (
	"github.com/grandcolline/clean-arch-demo/entity"
	"github.com/grandcolline/clean-arch-demo/usecase"
)

type userController struct {
	userService usecase.UserService
}

type UserController interface {
	CreateUser(user *entity.User) error
	GetUsers() ([]*entity.User, error)
}

func NewUserController(us usecase.UserService) UserController {
	return &userController{us}
}

func (userController *userController) CreateUser(user *entity.User) error {
	// 内側の処理のための技術的な関心事を記載
	return userController.userService.Create(user)
}

func (userController *userController) GetUsers() ([]*entity.User, error) {
	u := []*entity.User{}
	us, err := userController.userService.Get(u)
	if err != nil {
		return nil, err
	}
	return us, nil
}
