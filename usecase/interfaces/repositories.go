package interfaces

import "github.com/grandcolline/clean-arch-demo/entity"

type UserRepository interface {
	Store(entity.User) (int, error)
	FindByName(string) ([]entity.User, error)
	FindAll() ([]entity.User, error)
}
