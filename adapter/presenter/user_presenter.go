package presenter

import (
	"github.com/grandcolline/clean-arch-demo/entity"
)

type userPresenter struct {
}

func NewUserPresenter() UserPresenter {
	return &userPresenter{}
}

type UserPresenter interface {
	ResponseUsers(us []*entity.User) []*entity.User
}

func (userPresenter *userPresenter) ResponseUsers(us []*entity.User) []*entity.User {
	for _, u := range us {
		u.LastName = u.LastName + "さま"
	}
	return us
}
