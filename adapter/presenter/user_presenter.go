package presenter

import (
	"fmt"
	"net/http"

	"github.com/grandcolline/clean-arch-demo/entity"
)

// UserPresenter ユーザプレゼンタ
type UserPresenter struct {
	Writer http.ResponseWriter
}

// func NewUserPresenter(w http.ResponseWriter) usecase.UserOutputPort {
// 	return &UserPresenter{
// 		w: w,
// 	}
// }

// Render 名前に「さま」をつけて表示する
func (p *UserPresenter) Render(user *entity.User) error {
	fmt.Fprint(p.Writer, user.Name+"さま")
	return nil
}
