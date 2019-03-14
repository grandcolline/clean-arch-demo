package model

import (
	"github.com/grandcolline/clean-arch-demo/entity"
	"github.com/jinzhu/gorm"
)

// User ユーザモデルのデータ構造
type User struct {
	gorm.Model
	Name  string `gorm:"size:20;not null"`
	Email string `gorm:"size:100;not null"`
}

// ToEntity ユーザモデルをユーザエンティティに詰め直します。
func (m *User) ToEntity() (e entity.User) {
	e.ID = uint32(m.ID)
	e.Name = m.Name
	e.Email = m.Email
	return
}
