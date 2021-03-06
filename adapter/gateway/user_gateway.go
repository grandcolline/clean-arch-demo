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

// NewUserGateway ユーザゲートウェイを作成します。
func NewUserGateway(conn *gorm.DB) usecase.UserRepositoryPort {
	return &UserGateway{
		Conn: conn,
	}
}

// Store ユーザの新規追加します。
func (g *UserGateway) Store(u *entity.User) (*entity.User, error) {
	user := &model.User{
		Name:  u.Name,
		Email: u.Email,
	}

	// データをインサートする
	if err := g.Conn.Create(user).Error; err != nil {
		return nil, err
	}

	return user.ToEntity(), nil
}

// Update ユーザ情報を更新します。
func (g *UserGateway) Update(u *entity.User) error {
	user := &model.User{
		UUID:  u.UUID,
		Name:  u.Name,
		Email: u.Email,
	}

	// データをアップデートする
	return g.Conn.Omit("created_at").Save(user).Error
}

// Delete ユーザの削除をする
func (g *UserGateway) Delete(uuid string) error {
	user := &model.User{
		UUID: uuid,
	}

	// データを削除する
	return g.Conn.Delete(user).Error
}

/*
FindByName はユーザ名でユーザを検索します。

1件もヒットしなかった場合は、エラーでなく空のエンティティを返します。
*/
func (g *UserGateway) FindByName(name string) (*[]entity.User, error) {
	users := []model.User{}

	// データを取得する
	if err := g.Conn.Where("name = ?", name).Find(&users).Error; err != nil {
		return nil, err
	}

	// エンティティの作成
	e := make([]entity.User, len(users))
	for i, user := range users {
		e[i] = *user.ToEntity()
	}

	return &e, nil
}

/*
FindByID はIDでユーザを検索します。

1件もヒットしなかった場合は、エラーでなく空のエンティティを返します。
*/
func (g *UserGateway) FindByID(uuid string) (*entity.User, error) {
	user := model.User{
		UUID: uuid,
	}

	// DBからデータを取得する
	if err := g.Conn.First(&user).Error; err != nil {
		return nil, err
	}

	// エンティティの返却
	return user.ToEntity(), nil
}

// FindAll は全ユーザを検索します。
func (g *UserGateway) FindAll() (*[]entity.User, error) {
	users := []model.User{}

	// データを検索します
	if err := g.Conn.Find(&users).Error; err != nil {
		return nil, err
	}

	// エンティティの作成
	e := make([]entity.User, len(users))
	for i, user := range users {
		e[i] = *user.ToEntity()
	}

	return &e, nil
}

// IsNotFound はエラーがレコードが存在しなかったためかを判定します。
func (g *UserGateway) IsNotFound(err error) bool {
	return err == gorm.ErrRecordNotFound
}
