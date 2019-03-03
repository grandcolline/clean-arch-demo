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

// Store ユーザの新規追加をする
func (g *UserGateway) Store(u entity.User) (id uint32, err error) {
	user := &User{
		Name:  u.Name,
		Email: u.Email,
	}

	if err = g.Conn.Create(user).Error; err != nil {
		return
	}

	return uint32(user.ID), nil
}

// FindByName 名前でユーザを検索する
func (g *UserGateway) FindByName(name string) (d []entity.User, err error) {
	users := []User{}
	if err = g.Conn.Where("name = ?", name).Find(&users).Error; err != nil {
		return
	}

	// エンティティの作成
	n := len(users)
	d = make([]entity.User, n)
	for i := 0; i < n; i++ {
		d[i] = users[i].ToEntity()
	}

	return
}

// FindByID IDでユーザを検索する
func (g *UserGateway) FindByID(id uint32) (d entity.User, err error) {
	user := User{
		Model: gorm.Model{ID: uint(id)},
	}
	if err = g.Conn.First(&user).Error; err != nil {
		return
	}

	// エンティティの作成
	d = user.ToEntity()

	return
}

// FindAll 全ユーザを検索する
func (g *UserGateway) FindAll() (d []entity.User, err error) {
	users := []User{}
	if err = g.Conn.Find(&users).Error; err != nil {
		return
	}

	// エンティティの作成
	n := len(users)
	d = make([]entity.User, n)
	for i := 0; i < n; i++ {
		d[i] = users[i].ToEntity()
	}

	return
}
