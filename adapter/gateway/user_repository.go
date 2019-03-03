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
	d = usersToEntities(users)

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
	d = userToEntity(user)

	return
}

// FindAll 全ユーザを検索する
func (g *UserGateway) FindAll() (d []entity.User, err error) {
	users := []User{}
	if err = g.Conn.Find(&users).Error; err != nil {
		return
	}

	// エンティティの作成
	d = usersToEntities(users)

	return
}

// userToEntity ユーザをユーザエンティティに詰め直す
func userToEntity(u User) (e entity.User) {
	e.ID = uint32(u.ID)
	e.Name = u.Name
	e.Email = u.Email
	return
}

// usersToUserEntities ユーザのスライスをユーザエンティティのスライスに詰め直す
func usersToEntities(us []User) (es []entity.User) {
	n := len(us)
	es = make([]entity.User, n)
	for i := 0; i < n; i++ {
		es[i].ID = uint32(us[i].ID)
		es[i].Name = us[i].Name
		es[i].Email = us[i].Email
	}
	return
}
