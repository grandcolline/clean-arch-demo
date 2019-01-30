package inputport

import "github.com/grandcolline/clean-arch-demo/entity"

type UserInputport interface {
	Store(user *entity.User) error
	FindAll(users []*entity.User) ([]*entity.User, error)
}
