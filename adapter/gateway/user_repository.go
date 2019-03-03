package gateway

import (
	"github.com/grandcolline/clean-arch-demo/entity"
	"github.com/grandcolline/clean-arch-demo/usecase"
	"github.com/jinzhu/gorm"
)

// UserGateway ユーザゲートウェイ
type UserGateway struct {
	Conn *gorm.DB
}

// NewUserGateway ユーザゲートウェイの作成
func NewUserGateway(conn *gorm.DB) usecase.UserRepositoryPort {
	return &UserGateway{
		Conn: conn,
	}
}

// User ユーザテーブルのデータ構造
type User struct {
	gorm.Model
	Name  string `gorm:"size:20;not null"`
	Email string `gorm:"size:100;not null"`
}

// func (r *UserGateway) Store(u entity.User) (id int, err error) {
// 	user := &User{
// 		Name:  u.Name,
// 		Email: u.Email,
// 	}
//
// 	if err = r.Conn.Create(user).Error; err != nil {
// 		return
// 	}
//
// 	return int(user.ID), nil
// }

// FindByName 名前でユーザを検索する
func (r *UserGateway) FindByName(name string) (d []entity.User, err error) {
	users := []User{}
	if err = r.Conn.Where("name = ?", name).Find(&users).Error; err != nil {
		return
	}

	n := len(users)
	d = make([]entity.User, n)
	for i := 0; i < n; i++ {
		d[i].ID = uint32(users[i].ID)
		d[i].Name = users[i].Name
		d[i].Email = users[i].Email
	}
	return
}

// FindByID IDでユーザを検索する
func (r *UserGateway) FindByID(id uint32) (d entity.User, err error) {
	user := User{
		Model: gorm.Model{ID: uint(id)},
	}
	if err = r.Conn.First(&user).Error; err != nil {
		return
	}

	// エンティティの作成
	d.ID = uint32(user.ID)
	d.Name = user.Name
	d.Email = user.Email

	return
}

// FindAll 全ユーザを検索する
func (r *UserGateway) FindAll() (d []entity.User, err error) {
	users := []User{}
	if err = r.Conn.Find(&users).Error; err != nil {
		return
	}

	n := len(users)
	d = make([]entity.User, n)
	for i := 0; i < n; i++ {
		d[i].ID = uint32(users[i].ID)
		d[i].Name = users[i].Name
		d[i].Email = users[i].Email
	}
	return
}
