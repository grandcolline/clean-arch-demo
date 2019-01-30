package outputport

import "github.com/grandcolline/clean-arch-demo/entity"

type UserOutputport interface {
	ResponseUsers(es []*entity.User) []*entity.User
}
