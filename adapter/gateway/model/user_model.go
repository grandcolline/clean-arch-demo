package model

import (
	"github.com/grandcolline/clean-arch-demo/entity"
	"github.com/jinzhu/gorm"
)

// User ユーザモデルのデータ構造
type User struct {
	gorm.Model
	UUID  string `gorm:"size:20;not null;primary_key"`
	Name  string `gorm:"size:20;not null"`
	Email string `gorm:"size:100;not null"`
}

// ToEntity ユーザモデルをユーザエンティティに詰め直します。
func (m *User) ToEntity() *entity.User {
	return &entity.User{
		UUID:  m.UUID,
		Name:  m.Name,
		Email: m.Email,
	}
}
