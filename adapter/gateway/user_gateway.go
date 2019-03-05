package gateway

import (
	"github.com/grandcolline/clean-arch-demo/adapter/gateway/model"
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
func (g *UserGateway) Store(u entity.User) (e entity.User, err error) {
	user := &model.User{
		Name:  u.Name,
		Email: u.Email,
	}

	if err = g.Conn.Create(user).Error; err != nil {
		return
	}

	e = user.ToEntity()
	return
}

// Update ユーザ情報を更新する
func (g *UserGateway) Update(u entity.User) (err error) {
	user := &model.User{
		Model: gorm.Model{ID: uint(u.ID)},
		Name:  u.Name,
		Email: u.Email,
	}

	if err = g.Conn.Omit("created_at").Save(user).Error; err != nil {
		return
	}

	return
}

// Delete ユーザの削除をする
func (g *UserGateway) Delete(u entity.User) (err error) {
	user := &model.User{
		Model: gorm.Model{ID: uint(u.ID)},
		Name:  u.Name,
		Email: u.Email,
	}

	if err = g.Conn.Delete(user).Error; err != nil {
		return
	}

	return
}

// FindByName 名前でユーザを検索する
func (g *UserGateway) FindByName(name string) (d []entity.User, err error) {
	users := []model.User{}
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
	user := model.User{
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
	users := []model.User{}
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
