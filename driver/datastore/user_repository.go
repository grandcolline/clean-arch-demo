package datastore

import (
	"fmt"

	"github.com/grandcolline/clean-arch-demo/entity"
	"github.com/jinzhu/gorm"
)

type userRepository struct {
	db *gorm.DB
}

type UserRepository interface {
	Store(user *entity.User) error
	FindAll(users []*entity.User) ([]*entity.User, error)
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db}
}

func (userRepository *userRepository) Store(user *entity.User) error {
	return userRepository.db.Save(user).Error

}

func (userRepository *userRepository) FindAll(users []*entity.User) ([]*entity.User, error) {

	err := userRepository.db.Find(&users).Error
	if err != nil {
		return nil, fmt.Errorf("sql error", err)
	}

	return users, nil
}
